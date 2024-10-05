package dto

type DecodeRequest struct {
	InputString string `json:"InputString"`
}

type DecodeResponse struct {
	OutputString string `json:"OutputString"`
}