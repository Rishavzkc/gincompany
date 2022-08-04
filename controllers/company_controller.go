package controllers

import (
	"final/repositories"
	"final/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	repo    repositories.CompanyRepository = repositories.CompanyRepository{}
	service services.CompanyService        = *services.NewCompanyService(&repo)
)

func Show(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})
		return
	}

	result, err := service.ShowCompany(int(newid))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func ShowAll(c *gin.Context) {

	result, err := service.ShowAllCompanies()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

func Create(c *gin.Context) {
	var dto services.CompanyDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	result, err := service.CreateCompany(dto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}
	c.JSON(200, result)
}

func Update(c *gin.Context) {
	var dto services.CompanyDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateCompany(dto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "ID has to be integer"})

		return
	}

	result, err := service.DeleteCompany(int(newid))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	c.JSON(200, result)
}

//	repo := &repositories.CompanyRepository{}
//	service := services.NewCompanyService(repo)
