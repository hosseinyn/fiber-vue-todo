package models

type Todo struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Title        string    `json:"title"`
	Done         bool      `json:"done"`
	TargetUserID uint      `gorm:"default:null"`
	TargetUser   UserModel `gorm:"foreignKey:TargetUserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
