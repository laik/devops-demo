package user

import (
	"github.com/laik/devops-demo/pkg/service"
)
// {"username":"abc","age":123}

type User struct {
	i service.IService
}

func NewUser(i service.IService) *User {
	return &User{i}
}
