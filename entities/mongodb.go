package entities

type File struct {
	DocumentId string `bson:"-"`
	Name       string `bson:"name"`
	Data       []byte `bson:"data"`
}
