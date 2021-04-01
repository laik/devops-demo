package service

import (
	"github.com/laik/devops-demo/pkg/core"
	"github.com/laik/devops-demo/pkg/store"
)

// Query
// RDBMS  where a=123 and b=xx
// KV	  (a=123,b=xx)|(a=321,b=123)

// Mock

type IService interface {
	Create(db, table string, obj core.IObject) error
	Update(db, table string, src core.IObject, target core.IObject) error
	Delete(db, table string, uuid string) error
	Query(db, table string, c store.Condition) ([]core.IObject, error)
	QueryOne(db, table string, c store.Condition) (core.IObject, error)
	Range(db, table string, c store.Condition, f func(core.IObject) error) error
}

type BaseService struct {
	store.IStore
}

func (b BaseService) Create(db, table string, obj core.IObject) error {
	panic("implement me")
}

func (b BaseService) Update(db, table string, src core.IObject, target core.IObject) error {
	panic("implement me")
}

func (b BaseService) Delete(db, table string, uuid string) error {
	panic("implement me")
}

func (b BaseService) Query(db, table string, c store.Condition) ([]core.IObject, error) {
	panic("implement me")
}

func (b BaseService) QueryOne(db, table string, c store.Condition) (core.IObject, error) {
	panic("implement me")
}

func (b BaseService) Range(db, table string, c store.Condition, f func(core.IObject) error) error {
	panic("implement me")
}

func NewBaseService(s store.IStore) IService {
	return &BaseService{s}
}
