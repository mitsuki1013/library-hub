package model

type Company struct {
	ID   string `gorm:"id"`
	Name string `gorm:"name"`
}
