package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsouza/go-gin-boilerplate/models"
)

type CompanyController struct{}

var companyModel = new(models.User)

func (u CompanyController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := companyModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(400, gin.H{"message": "bad request"})
	c.Abort()
	return
}
