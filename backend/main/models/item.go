package models

type Item struct {
	Id          int     `json:"id" gorm:"primaryKey; autoIncrement"`
	Name        string  `json:"itemName"`
	Price       float64 `json:"price"`
	PicUrl      string  `json:"picUrl"`
	Description string  `json:"description"`
	//CreateTime  time.Time `json:"createTime"`
	//UpdateTime  time.Time `json:"updateTime"`
}
