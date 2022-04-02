package vo

import (
	"time"
)

type OrderHistoryVo struct {
	OrderId    string    `json:"orderId"`
	TotalPrice float64   `json:"totalPrice"`
	OrderDate  time.Time `json:"orderDate"`
	ItemVoList []ItemVo  `json:"itemList"`
}
