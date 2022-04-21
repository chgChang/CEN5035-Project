package pojo

type Cart struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	ItemId   int    `json:"itemId"`
	Quantity int    `json:"quantity"`
}
