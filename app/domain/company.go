package domain

type CompanyRepository interface {
	Create(Company) (Company, error)
}

type Company struct {
	ID   string
	Name string
}
