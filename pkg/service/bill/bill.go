package bill

import (
	"github.com/laik/devops-demo/pkg/common"
	"github.com/laik/devops-demo/pkg/module/order"
	"github.com/laik/devops-demo/pkg/service"
	"github.com/laik/devops-demo/pkg/store"
)

type Bill struct {
	service.IService
}

// {"bill_id":"123","user_name":"ljm"}
func (b *Bill) BillByUser(name string) ([]order.UserBill, error) {
	data, err := b.Query(common.DB, common.Bill, store.Condition{"user_name": name})
	if err != nil {
		return nil, err
	}
	_ = data
	return nil, nil
}

func NewBill(i service.IService) *Bill {
	return &Bill{i}
}
