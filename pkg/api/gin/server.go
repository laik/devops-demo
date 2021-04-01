package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/laik/devops-demo/pkg/api"
	"net/http"
)

const GROUP = "v1"

var _ api.IServer = &Gin{}

type Gin struct {
	*gin.Engine
	addr string
}

func NewGinServer() *Gin {
	route := gin.Default()
	return &Gin{Engine: route}
}

func (g *Gin) Run() error {
	return g.Engine.Run(g.addr)
}

func (g *Gin) Listen(addr string) api.IServer {
	g.addr = addr
	return g
}

func (g *Gin) Register(s string, handler http.Handler) error {
	panic("implement me")
}
