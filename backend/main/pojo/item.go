package pojo

type Item struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	PicUrl      string  `json:"pic_url"`
	Description string  `json:"description"`
}
