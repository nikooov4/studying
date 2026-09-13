package service

import "niki/internal/repository"

type Service struct {
	Repo *repository.Repository
}

func (s *Service) Service() string {
	ss := s.Repo.Return()
	return "service -> " + ss
}
