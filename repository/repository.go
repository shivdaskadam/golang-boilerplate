package repository

import (
	svc "github.com/shivdaskadam/golang-boilerplate/iface"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (svc.Repository, error) {
	return &repository{
		db: db,
	}, nil
}

func (r *repository) DB() *gorm.DB {
	return r.db
}
