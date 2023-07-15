package dto

type Request struct {
	RequestId  string
	DocumentId string
	Filename   string
}

type HeaderStorePhotoProfile struct {
	RequestId string `header:"request_id"`
	Filename  string `header:"filename"`
}

type Header struct {
	RequestId  string `header:"request_id"`
	DocumentId string `header:"document_id"`
}
