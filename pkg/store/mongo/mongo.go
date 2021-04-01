package mongo

import (
	"github.com/laik/devops-demo/pkg/core"
	"github.com/laik/devops-demo/pkg/store"
)

type Mongo struct {
}

func (m Mongo) Apply(database, table string, obj core.IObject) error {
	panic("implement me")
}

func (m Mongo) Get(database, table string, c store.Condition) (core.IObject, error) {
	panic("implement me")
}

func (m Mongo) Del(database, table string, c store.Condition) error {
	panic("implement me")
}

func NewMongo() store.IStore {
	return &Mongo{}
}
