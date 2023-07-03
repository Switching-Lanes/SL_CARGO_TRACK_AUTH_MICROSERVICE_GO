package routes

import (
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/handlers"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты приложения
func SetupRouter() *gin.Engine {
	router := gin.New()

	// Маршруты для регистрации и аутентификации

	ShipperGroup := router.Group("/shipper")
	// ShipperGroup.Use(middleware.AuthShipperMiddleware(), middleware.AuthorizeShipper("shipper"))
	{
		ShipperGroup.POST("/register", handlers.ShipperRegisterHandler)
		ShipperGroup.POST("/login", handlers.ShipperLogin)
		ShipperGroup.GET("/confirm-email/:confirmationLink", handlers.ConfirmEmailGetHandler)
		ShipperGroup.POST("/confirm-email", handlers.ConfirmEmailHandler)
	}

	CompanyGroup := router.Group("/company")
	{
		CompanyGroup.POST("/register", handlers.FreightCompanyRegisterHandler)
		CompanyGroup.POST("/login", handlers.FreightCompanyLogin)
		CompanyGroup.POST("/register/employee", middleware.AuthFreightCompanyMiddleware(), handlers.EmployeeRegisterHandler)
	}

	AdminGroup := router.Group("admin/")
	{
		AdminGroup.GET("/confirm-register-company/:confirmationLink", handlers.ConfirmRegisterFreightCompany)
	}

	RefreshTokenGroup := router.Group("/refresh_token")
	{
		RefreshTokenGroup.POST("/freight_company", handlers.FreightCompanyRefreshTokenHandler)
		// RefreshTokenGroup.POST("/shipper", handlers.ShipperRefreshTokenHandler)
		// RefreshTokenGroup.POST("/freight_company_employee", handlers.FreightCompanyEmployeeRefreshTokenHandler)

	}

	return router
}
