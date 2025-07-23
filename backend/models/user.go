package models

type UserModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
