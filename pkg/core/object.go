package core

import "encoding/json"

/*
{
	"kind": "action",
	"version": "20210303120901",
	"spec": { "data_list":[1,2,3] }
}
*/
type Kind = string

type IObject interface {
	UUID() string
	GetKind() Kind
	Get(string) interface{}
	Set(string, interface{})
}

//var _ IObject = &Object{}

// Go 3种类型
// 1. 指针类型  *Struct                               --> new(),&struct{}
// 2. 引用类型   []string,map[string]interface{},chan --> make()
// 3. 值类型    int...,string...                     -->  x:=1

type Object map[string]interface{}

func (o Object) GetKind() Kind {
	panic("implement me")
}

func FromBytes(b []byte) (Object, error) {
	o := make(Object)
	if err := json.Unmarshal(b, &o); err != nil {
		return nil, err // nil == 0x00
	}
	return o, nil
}

func Iter(n int64) []struct{} { return make([]struct{}, 0, n) }

// Sized 可被计算长度都可以放在stack
// 内存逃逸
type None struct{} // 0x00

func example() {
	none := &None{}
	_ = none

	channel := make(chan struct{})

	channel <- struct{}{}

	for _, v := range Iter(100) {
		_ = v
	}
}
