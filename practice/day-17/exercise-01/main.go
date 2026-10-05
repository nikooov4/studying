package main

import (
	"io"
	"encoding/json"
)

type CreateTaskRequest struct {
	Title string `json:"title"`
}

func DecodeCreateTask(r io.Reader) (CreateTaskRequest, error) {
	d := json.NewDecoder(r)
	task := CreateTaskRequest{}
	err := d.Decode(&task)
	

	return task, err
}
