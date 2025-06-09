package model

import "github.com/relvacode/iso8601"

type Document struct {
	Filename string `json:"filename"`
	Author   string `json:"author"`
	Cells    []Cell `json:"cells"`
}

type Cell struct {
	Input      string       `json:"input"`
	InputTime  iso8601.Time `json:"input_time"`
	Output     string       `json:"output"`
	OutputTime iso8601.Time `json:"output_time"`
	ExecTime   string       `json:"exec_time"`
}
