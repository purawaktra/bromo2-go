package dto

type BaseResponse struct {
	ResponseId   string `bson:"response_id"`
	RequestId    string `bson:"request_id"`
	Success      bool   `bson:"success"`
	ResponseTime string `bson:"response_time"`
}

type ResponseError struct {
	BaseResponse
	Errors any `bson:"errors,omitempty"`
}

type ResponseSuccess struct {
	BaseResponse
	Message string `bson:"message"`
	Data    any    `bson:"data"`
}
