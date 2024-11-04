package usecase

import "library-hub/app/domain"

type CreateCompanyUseCase interface {
	Execute(CreateCompanyInput) (CreateCompanyOutput, error)
}

type CreateCompanyInput struct {
	ID   string
	Name string
}

type CreateCompanyOutput struct {
	ID   string
	Name string
}

type createCompanyInteractor struct {
	repo domain.CompanyRepository
}

func NewCreateCompanyInteractor(
	repo domain.CompanyRepository,
) CreateCompanyUseCase {
	return createCompanyInteractor{
		repo: repo,
	}
}

func (c createCompanyInteractor) Execute(input CreateCompanyInput) (CreateCompanyOutput, error) {
	output, err := c.repo.Create(domain.Company{
		ID:   input.ID,
		Name: input.Name,
	})
	if err != nil {
		return CreateCompanyOutput{}, err
	}

	return CreateCompanyOutput{
		ID:   output.ID,
		Name: output.Name,
	}, nil
}
