package dept

import (
	"github.com/laik/devops-demo/pkg/service"
)

type Dept struct {
	service.IService
}

func NewDept(s *service.BaseService) *Dept {
	return &Dept{s}
}
