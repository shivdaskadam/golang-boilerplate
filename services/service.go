package services

import (
	svc "github.com/shivdaskadam/golang-boilerplate/iface"
)

type service struct {
	repository svc.Repository
}

func NewService(repo svc.Repository) svc.Service {
	return &service{
		repository: repo,
	}
}
