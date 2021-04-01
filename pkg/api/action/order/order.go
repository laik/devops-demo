package order

import (
	"github.com/gin-gonic/gin"
	gin2 "github.com/laik/devops-demo/pkg/api/gin"
	service_bill "github.com/laik/devops-demo/pkg/service/bill"
)

type Order struct {
	*gin2.Gin
	bill *service_bill.Bill
}

// /action/:name
func (o *Order) QueryBillByName(g *gin.Context) {
	o.bill.BillByUser(g.Param("name"))
	//g.JSON(http.StatusOK, data)
}
