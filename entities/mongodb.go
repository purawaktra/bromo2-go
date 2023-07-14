package entities

type File struct {
	Name string `bson:"name"`
	Data []byte `bson:"data"`
}
