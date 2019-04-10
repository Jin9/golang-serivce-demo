package model

// Customer is representation of a customer
type Customer struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// MessageResponse is representation of a status response
type MessageResponse struct {
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

// TokenResponse is represent of a insert response message
type TokenResponse struct {
	Token string `json:"token"`
}

// MessageInsertResponse is represent a model which insert success
type MessageInsertResponse struct {
	StatusCode    string        `json:"statusCode"`
	StatusMessage string        `json:"statusMessage"`
	TokenResponse TokenResponse `json:"tokenResponse"`
}

// MessageCustomerRequest is represent a model of FindCustomer request
type MessageCustomerRequest struct {
	Token string `json:"token"`
}

// MessageCustomerResponse is represent a model which find success
type MessageCustomerResponse struct {
	StatusCode    string    `json:"statusCode"`
	StatusMessage string    `json:"statusMessage"`
	Customer      *Customer `json:"customer"`
}

// MessageListCustomersResponse is represent a model of list of customer detail
type MessageListCustomersResponse struct {
	StatusCode    string      `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	Customers     []*Customer `json:"customers"`
}

// MessageUpdateCustomerRequest is represent a model of update request
type MessageUpdateCustomerRequest struct {
	Token    string    `json:"token"`
	Customer *Customer `json:"customer"`
}

// MessageUpdateCustomerResponse is represent a model which insert success
type MessageUpdateCustomerResponse struct {
	StatusCode    string        `json:"statusCode"`
	StatusMessage string        `json:"statusMessage"`
	TokenResponse TokenResponse `json:"tokenResponse"`
}

// NewCustomer is a function to create new model of Customer
func NewCustomer(name string, age int, email string, phone string) *Customer {
	return &Customer{
		Name:  name,
		Age:   age,
		Email: email,
		Phone: phone,
	}
}

// NewMessageResponse is created an error message response
func NewMessageResponse(statusCode string, statusMessage string) *MessageResponse {
	return &MessageResponse{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
	}
}

// NewMessageInsertResponse is created a success insert response
func NewMessageInsertResponse(statusCode string, statusMessage string, token string) *MessageInsertResponse {
	return &MessageInsertResponse{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		TokenResponse: TokenResponse{Token: token},
	}
}

// NewMessageCustomerResponse is created a success query customer by token response
func NewMessageCustomerResponse(statusCode string, statusMessage string, customer *Customer) *MessageCustomerResponse {
	return &MessageCustomerResponse{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		Customer:      customer,
	}
}

// NewMessageListCustomersResponse is created a success query all customer response
func NewMessageListCustomersResponse(statusCode string, statusMessage string, customers []*Customer) *MessageListCustomersResponse {
	return &MessageListCustomersResponse{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		Customers:     customers,
	}
}

// NewMessageUpdateCustomerResponse is create a success update customer response
func NewMessageUpdateCustomerResponse(statusCode string, statusMessage string, token string) *MessageUpdateCustomerResponse {
	return &MessageUpdateCustomerResponse{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
		TokenResponse: TokenResponse{Token: token},
	}
}
