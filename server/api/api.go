package api

import (
	"termbook/model"
	"time"
)

type NewDocumentRequest struct {
	Filename string `json:"filename"`
}

type OpenDocumentRequest struct {
	Filename string `json:"filename"`
}

type SaveDocumentRequest struct {
	model.Document
}

type ListDocumentResponse struct {
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type RunCellRequest struct {
	Input string `json:"input" example:"echo Hello World"`
}

type RunCellResponse struct {
	Output     string    `json:"output"`
	OutputTime time.Time `json:"output_time" example:"2025-06-09T13:11:45.3390718+08:00"`
	ExecTime   string    `json:"exec_time"   example:"150ms"`
}
