package models

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
