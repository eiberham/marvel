package model

type ResponseError struct {
	Code int `json:"code"`
	Status string `json:"status"`
}
