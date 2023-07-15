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
	RetrieveFileByDocumentId(documentId string) (entities.File, error)
	StoreFile(filename string, data []byte) (Data, error)
	RemoveFileByDocumentId(documentId string) error
}

func (uc Bromo2Usecase) RetrieveFileByDocumentId(documentId string) (entities.File, error) {
	// convert input to entity
	doc := entities.File{DocumentId: documentId}

	// call repo for the doc
	docs, err := uc.repo.RetrieveFileByDocumentId(doc)

	// check for error on call repo
	if err != nil {
		utils.Error(errors.New("error call usecase"), "RetrieveFileByDocumentId", doc)
		return entities.File{}, err
	}

	// convert entity to dto
	result := entities.File{
		Name: docs.Name,
		Data: docs.Data,
	}

	// create return
	return result, nil
}

func (uc Bromo2Usecase) StoreFile(filename string, data []byte) (entities.File, error) {
	// check for file name if empty
	if filename == "" {
		utils.Error(errors.New("filename can not be empty"), "StoreFile", "")
		return entities.File{}, errors.New("filename can not be empty")
	}

	// convert input to entity
	doc := entities.File{
		Name: filename,
		Data: data,
	}

	// call repo for the doc and check err
	document, err := uc.repo.StoreFile(doc)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		return entities.File{}, err
	}

	// convert entity to dto
	result := entities.File{
		DocumentId: document.DocumentId,
	}

	// create return
	return result, nil
}

func (uc Bromo2Usecase) RemoveFileByDocumentId(documentId string) error {
	// skip checking a document documentId
	// convert input to entity
	doc := entities.File{DocumentId: documentId}

	// call repo for the doc
	err := uc.repo.RemoveFileByDocumentId(doc)

	// check for error on call repo
	if err != nil {
		utils.Error(errors.New("error call usecase"), "RemoveFileByDocumentId", doc)
		return err
	}

	// create return
	return nil
}
