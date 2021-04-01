package store

import "github.com/laik/devops-demo/pkg/core"

type Condition map[string]interface{}

// DAO Database Access Object
// RDBMS database.Sql
// Kv

type IStore interface {
	Apply(database, table string, obj core.IObject) error // 有就更新，没有就添加
	Get(database, table string, c Condition) (core.IObject, error)
	Del(database, table string, c Condition) error
}

type MemStore struct {
	data map[string]interface{}
}

func (m MemStore) Apply(database, table string, obj core.IObject) error {
	panic("implement me")
}

func (m MemStore) Get(database, table string, c Condition) (core.IObject, error) {
	panic("implement me")
}

func (m MemStore) Del(database, table string, c Condition) error {
	panic("implement me")
}

func NewMemStore() IStore {
	return &MemStore{data: make(map[string]interface{})}
}
