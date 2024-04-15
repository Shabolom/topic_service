package models

type RespError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
