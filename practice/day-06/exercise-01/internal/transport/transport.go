package transport

import (
	"fmt"

	"niki/internal/service"
)

func Trans(s *service.Service) {

	tt := s.Service()
	fmt.Println(tt)
}
