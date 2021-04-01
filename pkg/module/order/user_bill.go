package order

import "time"

type UserBill struct {
	UserName  string    `json:"user_name"`
	BillId    int       `json:"bill_id"`
	CreatedAt time.Time `json:"created_at"`
}
