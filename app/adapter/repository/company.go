package repository

import (
	"gorm.io/gorm"
	"library-hub/app/domain"
)

type CompanyMySQL struct {
	db *gorm.DB
}

func NewCompanyMySQL(db *gorm.DB) *CompanyMySQL {
	return &CompanyMySQL{db: db}
}

func (c CompanyMySQL) Create(company domain.Company) (domain.Company, error) {
	if err := c.db.Create(&company).Error; err != nil {
		return domain.Company{}, err
	}
	return company, nil
}
