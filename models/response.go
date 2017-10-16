package models

type Response struct {
	Errors []Error `json:"errors"`
	Data interface{} `json:"data"`
}

type Error struct {
	Message string `json:"message"`
	Code string `json:"code"`
}