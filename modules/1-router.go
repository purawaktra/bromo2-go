package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/bromo2-go/utils"
)

type Bromo2Router struct {
	engine *gin.Engine
	rh     Bromo2RequestHandler
}

func CreateBromo2Router(engine *gin.Engine, rh Bromo2RequestHandler) Bromo2Router {
	return Bromo2Router{
		engine: engine,
		rh:     rh,
	}
}

func (r Bromo2Router) Init(path string) {
	pathGroup := r.engine.Group(path, utils.CheckBasicAuth)
	pathGroup.GET("/select/file", r.rh.RetrieveFile)
	pathGroup.POST("/insert/file", r.rh.StoreFile)
	pathGroup.DELETE("/delete/file", r.rh.RemoveFile)

}
