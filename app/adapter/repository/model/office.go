package model

type Office struct {
	ID        string `gorm:"id"`
	Name      string `gorm:"name"`
	CompanyID string `gorm:"company_id"`
}
