package main

import (
	"fmt"
	"github.com/laik/devops-demo/pkg/api"
	"github.com/laik/devops-demo/pkg/api/action/user"
	"github.com/laik/devops-demo/pkg/api/gin"
	"github.com/laik/devops-demo/pkg/service"
	service_dept "github.com/laik/devops-demo/pkg/service/dept"
	service_user "github.com/laik/devops-demo/pkg/service/user"
	"github.com/laik/devops-demo/pkg/store"
)

// 需要提供一个用户查询CRUD
// 关键点： 后端数据库是什么类型
//         1. RDBMS --> Mysql,SqlServer,Oracle,......   Query Delete Update Create
//         2. KvStore ---> MongoDB,Neo4j,Etcd........   Put/Set Get Remove/Delete Update/Apply Range(fn())

func main() {
	fmt.Println("yce-api")

	istore := store.NewMemStore()
	baseService := service.NewBaseService(istore)

	ginServer := gin.NewGinServer()
	ginServer.Listen("0.0.0.0:8080")

	user.NewUser(ginServer, service_user.NewUser(baseService), service_dept.NewDept(baseService))
	//ginServer.Listen("127.0.0.1:8080")

	server := api.NewServers(
		ginServer,
	)
	for err := range server.RunAll() {
		panic(err)
	}
}
