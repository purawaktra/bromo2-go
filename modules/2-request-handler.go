package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/bromo2-go/dto"
	"github.com/purawaktra/bromo2-go/functions"
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
	RetrieveFile(c *gin.Context)
	StoreFile(c *gin.Context)
	RemoveFile(c *gin.Context)
}

func (rh Bromo2RequestHandler) RetrieveFile(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "RetrieveFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "RetrieveFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	request := dto.Request{
		RequestId:  header.RequestId,
		DocumentId: header.DocumentId,
	}
	utils.Debug("RetrieveFile", request)
	response, data, err := rh.ctrl.RetrieveFile(request)
	if err != nil {
		utils.Error(err, "RetrieveFile", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("response_id", response.ResponseId)
	c.Header("response_time", response.ResponseTime)
	c.Header("message", response.Message)
	c.Header("request_id", response.RequestId)
	c.Data(http.StatusOK, "", data)
	return
}

func (rh Bromo2RequestHandler) StoreFile(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get header and check for err
	var header dto.HeaderStorePhotoProfile
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "StoreFile", "")
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

	// call controller for the file and check err
	requestBody := dto.Request{
		RequestId: header.RequestId,
		Filename:  header.Filename,
	}

	utils.Debug("StoreFile", requestBody.Filename)
	response, err := rh.ctrl.StoreFile(requestBody, readData)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("response_id", response.ResponseId)
	c.Header("response_time", response.ResponseTime)
	c.Header("message", response.Message)
	c.Header("request_id", response.RequestId)
	c.Data(http.StatusOK, "", []byte(""))
	return
}

func (rh Bromo2RequestHandler) RemoveFile(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get header and check for err
	var header dto.Header
	err := c.BindHeader(&header)
	if err != nil {
		utils.Error(err, "RemoveFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// check for request id
	if functions.IsValidUUID(header.RequestId) {
		utils.Error(err, "RemoveFile", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// call controller for the profile picture and check err
	request := dto.Request{
		RequestId:  header.RequestId,
		DocumentId: header.DocumentId,
	}
	utils.Debug("RemoveFile", request)
	response, err := rh.ctrl.RemoveFile(request)
	if err != nil {
		utils.Error(err, "RemoveFile", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("response_id", response.ResponseId)
	c.Header("response_time", response.ResponseTime)
	c.Header("message", response.Message)
	c.Header("request_id", response.RequestId)
	c.Data(http.StatusOK, "", []byte(""))
	return
}
