package vo

type ItemVo struct {
	ItemId      int     `json:"itemId"`
	ItemName    string  `json:"itemName"`
	Price       float64 `json:"price"`
	PicUrl      string  `json:"picUrl"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}
