package pojo

type OrderItem struct {
	Email           string  `json:"email"`
	OrderNo         string  `json:"orderNo"`
	ItemId          int     `json:"itemId"`
	ItemName        string  `json:"itemName"`
	ItemImage       string  `json:"itemImage"`
	UnitPrice       float64 `json:"unitPrice"`
	ItemDescription string  `json:"itemDescription"`
	Quantity        int     `json:"quantity"`
	TotalPrice      float64 `json:"totalPrice"`
}
