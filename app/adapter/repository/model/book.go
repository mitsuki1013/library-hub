package model

type Book struct {
	ID       string `gorm:"id"`
	Name     string `gorm:"name"`
	OfficeID string `gorm:"office_id"`
	Isbn     string `gorm:"isbn"`
}
