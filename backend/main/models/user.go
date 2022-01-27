package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey; autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
}
