package Entity

import "bytes"

type Task struct {
	Label   string `json:"label"`
	Command string `json:"command"`
	Out bytes.Buffer
	Stderr bytes.Buffer
}