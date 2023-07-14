package modules

import (
	"errors"
	"github.com/purawaktra/bromo2-go/entities"
	"github.com/purawaktra/bromo2-go/utils"
)

type Bromo2Usecase struct {
	repo Bromo2Repo
}

func CreateBromo2Usecase(repo Bromo2Repo) Bromo2Usecase {
	return Bromo2Usecase{
		repo: repo,
	}
}

type Bromo2UsecaseInterface interface {
	StoreFile(filename string, data []byte) (Data, error)
}

func (uc Bromo2Usecase) StoreFile(filename string, data []byte) (Data, error) {
	// check for file name if empty
	if filename == "" {
		utils.Error(errors.New("filename can not be empty"), "StoreFile", "")
		return Data{}, errors.New("filename can not be empty")
	}

	// convert input to entity
	doc := entities.File{
		Name: filename,
		Data: data,
	}

	// call repo for the doc and check err
	documentId, err := uc.repo.StoreFile(doc)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		return Data{}, err
	}

	// convert entity to dto
	result := Data{
		DocumentId: documentId,
		Filename:   filename,
	}

	// create return
	return result, nil
}
