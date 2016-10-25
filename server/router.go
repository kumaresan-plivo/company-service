package server

import (
	"github.com/gin-gonic/gin"

	"../controllers"
	"../middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//health := new(controllers.HealthController)
	//router.GET("/health", health.Status)
	router.Use(middlewares.AuthMiddleware())
	v1 := router.Group("v1")
	{
		companyGroup := v1.Group("company")
		{
			company := new(controllers.CompanyController)
			companyGroup.GET("/:id", company.Retrieve)
			//userGroup.POST("/", company.Signup)
			//userGroup.DELETE("/:id", company.Delete)
			//userGroup.PUT("/:id", company.Update)
		}
	}
	return router

}
