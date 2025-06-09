package api

import "termbook/model"

type ResponseOpenDocument struct {
	Msg  string         `json:"msg"`
	Data model.Document `json:"data"`
}

type ResponseNewDocument struct {
	Msg  string         `json:"msg"`
	Data model.Document `json:"data"`
}

type ResponseBad struct {
	Msg string `json:"msg"`
}

type ResponseRunCell struct {
	Msg  string          `json:"msg"`
	Data RunCellResponse `json:"data"`
}
