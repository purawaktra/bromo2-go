package modules

import (
	"github.com/google/uuid"
	"github.com/purawaktra/bromo2-go/dto"
	"github.com/purawaktra/bromo2-go/utils"
	"strconv"
	"time"
)

type Bromo2Controller struct {
	uc Bromo2Usecase
}

func CreateBromo2Controller(uc Bromo2Usecase) Bromo2Controller {
	return Bromo2Controller{
		uc: uc,
	}
}

type Bromo2ControllerInterface interface {
	RetrieveFile(req dto.Request) (dto.Response, []byte, error)
	StoreFile(req dto.Request, data []byte) (any, error)
	RemoveFile(req dto.Request) (dto.Response, error)
}

func (ctrl Bromo2Controller) RetrieveFile(req dto.Request) (dto.Response, []byte, error) {
	// start timer
	start := time.Now()

	// call usecase for the document
	doc, err := ctrl.uc.RetrieveFileByDocumentId(req.DocumentId)

	// check for error on call usecase
	if err != nil {
		utils.Error(err, "RetrieveProfilePicture", req)

		end := time.Now()
		return dto.Response{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      false,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			Message:      err.Error(),
		}, nil, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.Response{
		ResponseId:   uuid.New().String(),
		RequestId:    req.RequestId,
		Success:      false,
		ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		Message:      "Success",
	}

	// create return
	return result, doc.Data, nil
}

func (ctrl Bromo2Controller) StoreFile(req dto.Request, data []byte) (dto.Response, error) {
	// start timer
	start := time.Now()

	// call usecase for the document and check err
	doc, err := ctrl.uc.StoreFile(req.Filename, data)
	if err != nil {
		utils.Error(err, "StoreFile", req.Filename)
		end := time.Now()
		return dto.Response{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      false,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			Message:      err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.Response{
		ResponseId:   uuid.New().String(),
		RequestId:    req.RequestId,
		Success:      false,
		ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		Message:      doc.DocumentId,
	}

	// create return
	return result, nil
}

func (ctrl Bromo2Controller) RemoveFile(req dto.Request) (dto.Response, error) {
	// start timer
	start := time.Now()

	// call usecase for the document and check err
	err := ctrl.uc.RemoveFileByDocumentId(req.DocumentId)
	if err != nil {
		utils.Error(err, "RemoveFile", req)

		end := time.Now()
		return dto.Response{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      false,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			Message:      err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.Response{
		ResponseId:   uuid.New().String(),
		RequestId:    req.RequestId,
		Success:      false,
		ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		Message:      "Success",
	}

	// create return
	return result, nil
}
