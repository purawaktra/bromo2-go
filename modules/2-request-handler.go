package modules

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/bromo2-go/dto"
	"github.com/purawaktra/bromo2-go/utils"
	"io"
	"net/http"
)

type Bromo2RequestHandler struct {
	ctrl Bromo2Controller
}

func CreateBromo2RequestHandler(ctrl Bromo2Controller) Bromo2RequestHandler {
	return Bromo2RequestHandler{
		ctrl: ctrl,
	}
}

type Bromo2RequestHandlerInterface interface {
	StoreFile(c *gin.Context)
}

func (rh Bromo2RequestHandler) StoreFile(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get a request id and check for err
	requestId := c.GetHeader("request_id")
	if requestId == "" {
		utils.Error(errors.New("request_id can not be empty"), "StoreFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// get a request id and check for err
	filename := c.GetHeader("filename")
	if filename == "" {
		utils.Error(errors.New("filename can not be empty"), "StoreFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// get request body and check err
	readData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse data to struct
	requestBody := dto.Request{
		RequestId: requestId,
		Filename:  filename,
		Data:      readData,
	}

	// call controller for the file and check err
	utils.Debug("StoreFile", requestBody)
	response, err := rh.ctrl.StoreFile(requestBody)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.JSON(http.StatusOK, response)
	return
}
