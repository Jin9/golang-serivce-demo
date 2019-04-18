package controller

import (
	"net/http"

	"gitlab.com/chinnawat.w/golang-service-demo/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init used for annouce route path of api
func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://172.20.10.10:8880"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType},
		AllowCredentials: true,
	}))

	// routing
	e.POST("/customer", service.InsertCustomer)
	e.POST("/customer/detail", service.FindCustomer)
	e.POST("/customer/delete", service.DeleteCustomer)
	e.GET("/customer", service.ListCustomer)
	e.PATCH("/customer", service.UpdateCustomer)

	return e
}
