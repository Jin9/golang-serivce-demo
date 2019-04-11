package service

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"../customer"
	"../model"
	"../storage"

	"github.com/labstack/echo"
)

func tokenGenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func prepareToken() string {
	for {
		token := tokenGenerator()
		checkToken := customer.CheckDuplicateToken(token)
		if checkToken != nil {
			return token
		}
	}
}

func prepareMessageResponse(statusCode string, statusMessage string) *model.MessageResponse {
	return model.NewMessageResponse(statusCode, statusMessage)
}

func prepareMessageInsertResponse(statusCode string, statusMessage string, token string) *model.MessageInsertResponse {
	return model.NewMessageInsertResponse(statusCode, statusMessage, token)
}

func prepareMessageFindResponse(statusCode string, statusMessage string, detail *model.Customer) *model.MessageCustomerResponse {
	return model.NewMessageCustomerResponse(statusCode, statusMessage, detail)
}

func prepareMessageAllListResponse(statusCode string, statusMessage string, customers []*model.CustomerDetail) *model.MessageListCustomersResponse {
	return model.NewMessageListCustomersResponse(statusCode, statusMessage, customers)
}

func prepareMessageUpdateCustomerResponse(statusCode string, statusMessage string, token string) *model.MessageUpdateCustomerResponse {
	return model.NewMessageUpdateCustomerResponse(statusCode, statusMessage, token)
}

func callInsertCustomer(cust *model.Customer) (token string, err error) {
	token = prepareToken()
	detail := model.NewCustomer(cust.Name, cust.Age, cust.Email, cust.Phone)
	if err = customer.InsertCustomer(token, detail); err != nil {
		return "", err
	}
	return token, nil
}

// InsertCustomer is used to save new customer detail
func InsertCustomer(c echo.Context) (err error) {
	cust := new(model.Customer)
	if err = c.Bind(&cust); err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	token, err := callInsertCustomer(cust)
	if err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	return c.JSON(http.StatusOK, prepareMessageInsertResponse("0", "success", token))
}

// FindCustomer is used to find customer detail by token
func FindCustomer(c echo.Context) (err error) {
	msgCustomerRequest := new(model.MessageCustomerRequest)
	if err = c.Bind(&msgCustomerRequest); err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	storage.SetUserToken(msgCustomerRequest.Token)
	detail, err := customer.FindCustomerDetailByToken(storage.GetUserToken())
	if err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	return c.JSON(http.StatusOK, prepareMessageFindResponse("0", "success", detail))
}

// ListCustomer is used show all customer detail
func ListCustomer(c echo.Context) (err error) {
	customers, err := customer.FindAllCustomerDetail()
	if err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	return c.JSON(http.StatusOK, prepareMessageAllListResponse("0", "success", customers))
}

func validateTokenForUpdate(token string) bool {
	return token == storage.GetUserToken()
}

func callUpdateCustomer(request *model.MessageUpdateCustomerRequest, c echo.Context) (err error) {
	err = customer.UpdateCustomerDetail(request.Token, request.Customer)
	if err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	newToken := prepareToken()
	err = customer.UpdateCustomerToken(request.Token, newToken)
	if err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	storage.SetUserToken("")
	return c.JSON(http.StatusOK, prepareMessageUpdateCustomerResponse("0", "success", newToken))
}

// UpdateCustomer is used for update customer detail
func UpdateCustomer(c echo.Context) (err error) {
	msgUpdateCustomerRequest := new(model.MessageUpdateCustomerRequest)
	if err = c.Bind(&msgUpdateCustomerRequest); err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	if !validateTokenForUpdate(msgUpdateCustomerRequest.Token) {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", "Invalid Token"))
	}

	return callUpdateCustomer(msgUpdateCustomerRequest, c)
}

func callDeleteCustomer(token string, c echo.Context) (err error) {
	if err = customer.DeleteCustomerByToken(token); err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	return c.JSON(http.StatusOK, prepareMessageResponse("0", "success"))
}

// DeleteCustomer is used for delete customer
func DeleteCustomer(c echo.Context) (err error) {
	msgCustomerRequest := new(model.MessageCustomerRequest)
	if err = c.Bind(&msgCustomerRequest); err != nil {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", err.Error()))
	}

	if !validateTokenForUpdate(msgCustomerRequest.Token) {
		return c.JSON(http.StatusOK, prepareMessageResponse("1003", "Invalid Token"))
	}

	return callDeleteCustomer(msgCustomerRequest.Token, c)
}
