package user

import (
	"github.com/gin-gonic/gin"
	gin2 "github.com/laik/devops-demo/pkg/api/gin"
	service_dept "github.com/laik/devops-demo/pkg/service/dept"
	service_user "github.com/laik/devops-demo/pkg/service/user"
)

type User struct {
	*gin2.Gin
	user *service_user.User
	dept *service_dept.Dept
}

// 依赖 注入
// /v1/user/:name -> {"username":"","age":12}  get one
// /v1/users -> [{"username":"","age":12}]	   list
// /v1/dept/:name -> {"dept_name":"","age":12} get one
// /v1/depts -> [{"dept_name":"","age":12}]	   list

func NewUser(g *gin2.Gin, user *service_user.User, dept *service_dept.Dept) *User {
	u := &User{Gin: g, user: user, dept: dept}
	group := u.Group(gin2.GROUP)
	group.GET("/user/:name", u.UserGetOne)
	group.GET("/users", u.UserList)
	group.GET("/dept/:name", u.DeptGetOne)
	group.GET("/depts", u.DeptList)

	return u
}

// /action/:name
func (u *User) UserGetOne(g *gin.Context) {
}

func (u *User) UserList(g *gin.Context) {
	//username := g.Param("name")
}

// /v1/action/:name -> {"username":"","age":12}  get one
func (u *User) DeptGetOne(g *gin.Context) {
	//username := g.Param("name")
}

// /v1/users -> [{"username":"","age":12}]	   list
func (u *User) DeptList(g *gin.Context) {
	//username := g.Param("name")
}
