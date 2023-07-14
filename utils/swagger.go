package utils

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
	"io"
	"os"
)

const docTemplate = `{
  "openapi": "3.0.0",
  "info": {
    "description": "This is default API documentation",
    "version": "1.0.0",
    "title": "Default API documentation",
    "contact": {
      "email": "akbar.muhammadakbarmaulana@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

type SwaggerRouter struct {
	engine *gin.Engine
}

func CreateSwaggerRouter(engine *gin.Engine) SwaggerRouter {
	return SwaggerRouter{
		engine: engine,
	}
}

func (r SwaggerRouter) InitRouter(path string) {
	pathGroup := r.engine.Group(path)
	pathGroup.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func init() {
	// open swagger file and check err
	filePath := "swagger.json"
	file, err := os.Open(filePath)
	if err != nil {
		Fatal(err, "init", "")
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			Fatal(err, "init", "")
			return
		}
	}()

	// read file and check err
	data, err := io.ReadAll(file)
	if err != nil {
		Fatal(err, "init", "")
		return
	}

	// convert to string and check if empty
	swaggerJSON := string(data)
	if swaggerJSON != "" {
		SwaggerInfo.SwaggerTemplate = swaggerJSON
	}

	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
