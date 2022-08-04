package models

type Company struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func NewCompany(name string, location string) *Company {
	return &Company{
		Name:     name,
		Location: location,
	}
}

func NewCompanyWithID(id int, name string, location string) *Company {
	return &Company{
		ID:       id,
		Name:     name,
		Location: location,
	}
}
