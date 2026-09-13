package main

import (
	"niki/internal/repository"
	"niki/internal/service"
	"niki/internal/transport"
)

func main() {

	r := &repository.Repository{}
	s := &service.Service{
		Repo: r,
	}
	transport.Trans(s)
}
