package etcd

import (
	"github.com/laik/devops-demo/pkg/core"
	"github.com/laik/devops-demo/pkg/store"
)

type Etcd struct{}

func (e Etcd) Apply(database, table string, obj core.IObject) error {
	panic("implement me")
}

func (e Etcd) Get(database, table string, c store.Condition) (core.IObject, error) {
	panic("implement me")
}

func (e Etcd) Del(database, table string, c store.Condition) error {
	panic("implement me")
}

func NewEtcd() store.IStore {
	return &Etcd{}
}
