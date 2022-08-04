package repositories

import (
	"final/database"
	"final/models"
	"fmt"
)

type CompanyRepository struct {
}

func (s *CompanyRepository) GetAllCompany() ([]models.Company, error) {
	db := database.GetDatabase()
	var c []models.Company
	err := db.Find(&c).Error
	if err != nil {
		return nil, fmt.Errorf("cannot find products: %v", err)
	}

	return c, err
}

func (s *CompanyRepository) CreateCompany(c models.Company) (models.Company, error) {
	db := database.GetDatabase()
	err := db.Create(&c).Error

	if err != nil {
		return models.Company{}, err
	}
return c, nil
}

func (s *CompanyRepository) GetCompanyById(id int) (models.Company, error) {
	db := database.GetDatabase()
	var c models.Company
	err := db.First(&c, id).Error

	if err != nil {
		return models.Company{}, fmt.Errorf("cannot find company by id: %v", err)
	}
	return c, nil
}

func (s *CompanyRepository) DeleteCompanyById(id int) (models.Company, error) {
	db := database.GetDatabase()
	n, err := s.GetCompanyById(id)
	if err != nil {
		return models.Company{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.Company{}, fmt.Errorf("cannot delete : %v", err)
	}

	return n, nil
}

func (s *CompanyRepository) UpdateCompanyById(c models.Company) error {
	db := database.GetDatabase()

	err := db.Save(&c).Error
	if err != nil {
		return fmt.Errorf("cannot update : %v", err)
	}

	return nil
}
