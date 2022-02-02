package pojo

import "time"

type Order struct {
	OrderNo    string    `json:"orderNo"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
}
