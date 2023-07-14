package modules

import (
	"context"
	"fmt"
	"github.com/purawaktra/bromo2-go/entities"
	"github.com/purawaktra/bromo2-go/utils"
	"go.mongodb.org/mongo-driver/bson"
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

func (br Bromo2Repo) StoreFile(query entities.File) (string, error) {
	// set up config
	coll := br.db.Database(utils.MongoDBDatabase).Collection(utils.CollectionFiles)

	// create a bson data format and check err
	parsedData, err := bson.Marshal(query)
	if err != nil {
		utils.Error(err, "StoreFile", "")
		return "", err
	}

	// push data to mongo db and check err
	doc, err := coll.InsertOne(context.TODO(), parsedData)
	if err != nil {
		utils.Error(err, "StoreProfilePicture", "")
		return "", err
	}

	// create return
	documentId := fmt.Sprintf("%v", doc.InsertedID)
	utils.Debug("StoreProfilePicture", fmt.Sprintf("Success store file with ID : %v", doc.InsertedID))
	return documentId, nil
}
