package dto

type Response struct {
	ResponseId   string `bson:"response_id"`
	RequestId    string `bson:"request_id"`
	Success      bool   `bson:"success"`
	ResponseTime string `bson:"response_time"`
	Message      string `bson:"message"`
}
