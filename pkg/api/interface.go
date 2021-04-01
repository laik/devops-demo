package api

import (
	"fmt"
	"net/http"
)

type IServers []IServer

func NewServers(servers ...IServer) IServers {
	is := make(IServers, 0)
	is = append(is, servers...)
	return servers
}

func (i IServers) RunAll() <-chan error {
	errors := make(chan error, len(i))
	for _, server := range i {
		_server := server
		go func(is IServer) {
			if err := _server.Run(); err != nil {
				errors <- err
			}
		}(_server)
	}
	return errors
}

type IServer interface {
	Run() error
	Listen(addr string) IServer
	Register(string, http.Handler) error
}

type MyInterface interface {
	Do()
}

// Go 的结构休 是不能传递引用的
var _ IServer = &FakeServer{} // new(FakeServer)

type FakeServer struct {
	MyInterface // 指针 0x00000016  stack -> heap  引用计数器
}

func (f *FakeServer) Listen(addr string) IServer {
	panic("implement me")
}

// /v1/users
// /v2/users
// /yce-api/v1/users
func (f *FakeServer) Register(s string, handler http.Handler) error {
	panic("implement me")
}

func (f *FakeServer) Run() error {
	fmt.Printf("fakeServer run")
	return nil
}

func startFakeServer(fk FakeServer) {
	fk.Run()
	fk.Do()
}
