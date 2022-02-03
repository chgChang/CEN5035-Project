package form

type CartUpdateForm struct {
	ItemId   int `json:"itemId"`
	Quantity int `json:"quantity"`
}
