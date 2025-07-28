package models

type UserModel struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
}

type ChangePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
