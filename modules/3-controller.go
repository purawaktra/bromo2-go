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
	StoreFile(req dto.Request) (any, error)
}

func (ctrl Bromo2Controller) StoreFile(req dto.Request) (any, error) {
	// start timer
	start := time.Now()

	// call usecase for the document and check err
	data, err := ctrl.uc.StoreFile(req.Filename, req.Data)
	if err != nil {
		utils.Error(err, "StoreFile", req.Filename)
		end := time.Now()
		return dto.ResponseError{
			BaseResponse: dto.BaseResponse{
				ResponseId:   uuid.New().String(),
				RequestId:    req.RequestId,
				Success:      false,
				ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
			},
			Errors: err.Error(),
		}, err
	}

	// convert dto to response
	end := time.Now()
	result := dto.ResponseSuccess{
		BaseResponse: dto.BaseResponse{
			ResponseId:   uuid.New().String(),
			RequestId:    req.RequestId,
			Success:      true,
			ResponseTime: strconv.FormatInt(end.Sub(start).Microseconds(), 10),
		},
		Message: "Success StoreFile",
		Data:    data,
	}

	// create return
	return result, nil
}
