package vo

type CartVo struct {
	ItemList   []ItemVo `json:"itemList"`
	TotalPrice float64  `json:"totalPrice"`
}
