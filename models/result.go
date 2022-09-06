package models

type Result struct {
	Error string `json:"error"`
	Data interface{} `json:"data"`
	Code int `json:"code"`

}
