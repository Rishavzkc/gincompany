package interfaces

import "final/models"

type CompanyRepository interface {
	GetAllCompany() ([]models.Company, error)
	CreateCompany(c models.Company) (models.Company, error)
	GetCompanyById(id int) (models.Company, error)
	DeleteCompanyById(id int) (models.Company, error)
	UpdateCompany(c models.Company) error
}
