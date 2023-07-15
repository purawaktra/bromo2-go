package modules

import (
	"context"
	"fmt"
	"github.com/purawaktra/bromo2-go/entities"
	"github.com/purawaktra/bromo2-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bromo2Repo struct {
	db *mongo.Client
}

func CreateBromo2Repo(db *mongo.Client) Bromo2Repo {
	return Bromo2Repo{
		db: db,
	}
}

type Bromo2RepoInterface interface {
	RetrieveFileByDocumentId(query entities.File) (entities.File, error)
	StoreFile(query entities.File) (entities.File, error)
	RemoveFileByDocumentId(query entities.File) error
}

func (br Bromo2Repo) RetrieveFileByDocumentId(query entities.File) (entities.File, error) {
	// set up config
	coll := br.db.Database(utils.MongoDBDatabase).Collection(utils.CollectionFiles)
	id, err := primitive.ObjectIDFromHex(query.DocumentId)
	if err != nil {
		utils.Error(err, "RetrieveFileByDocumentId", query.DocumentId)
		return entities.File{}, err
	}
	filter := bson.M{"_id": id}

	// get document from mongo db
	doc := coll.FindOne(context.TODO(), filter)

	// check if err
	err = doc.Err()
	if err != nil {
		utils.Error(err, "RetrieveFileByDocumentId", filter)
		return entities.File{}, err
	}

	// parse document to struct
	var result entities.File
	err = doc.Decode(&result)

	// check if err parse document
	if err != nil {
		utils.Error(err, "RetrieveFileByDocumentId", "")
		return entities.File{}, err
	}

	// create return
	return result, nil
}

func (br Bromo2Repo) StoreFile(query entities.File) (entities.File, error) {
	// set up config
	coll := br.db.Database(utils.MongoDBDatabase).Collection(utils.CollectionFiles)

	// create a bson data format and check err
	parsedData, err := bson.Marshal(query)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		return entities.File{}, err
	}

	// push data to mongo db and check err
	doc, err := coll.InsertOne(context.TODO(), parsedData)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		return entities.File{}, err
	}

	// create return
	documentId := fmt.Sprintf("%v", doc.InsertedID)
	documentId = documentId[10:]
	documentId = documentId[:24]
	utils.Debug("StoreProfilePicture", fmt.Sprintf("Success store file with ID : %v", doc.InsertedID))
	return entities.File{DocumentId: documentId}, nil
}

func (br Bromo2Repo) RemoveFileByDocumentId(query entities.File) error {
	// set up config
	coll := br.db.Database(utils.MongoDBDatabase).Collection(utils.CollectionFiles)
	id, err := primitive.ObjectIDFromHex(query.DocumentId)
	if err != nil {
		utils.Error(err, "RemoveFileByDocumentId", query.DocumentId)
		return err
	}
	filter := bson.M{"_id": id}

	// delete document on mongo db
	result, err := coll.DeleteOne(context.TODO(), filter)

	// check if err
	if err != nil {
		utils.Error(err, "RemoveFileByDocumentId", filter)
		return err
	}

	// warn if deleted two
	if result.DeletedCount > 1 {
		utils.Warn("RemoveFileByDocumentId", "deleted more than one document")
	}

	// create return
	return nil
}
