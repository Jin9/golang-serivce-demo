package controller

import (
	"../service"

	"github.com/labstack/echo"
)

// Init used for annouce route path of api
func Init() *echo.Echo {
	e := echo.New()

	// routing
	e.POST("/customer", service.InsertCustomer)
	e.POST("/customer/detail", service.FindCustomer)
	e.POST("/customer/delete", service.DeleteCustomer)
	e.GET("/customer", service.ListCustomer)
	e.PATCH("/customer", service.UpdateCustomer)

	return e
}
