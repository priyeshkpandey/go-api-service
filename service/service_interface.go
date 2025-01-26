package service

import (
	"net/http"
)

type OrderServiceInterface interface {
	GetAllOrders(writer http.ResponseWriter, request *http.Request)
	GetOrderById(writer http.ResponseWriter, request *http.Request)
	CreateOrder(writer http.ResponseWriter, request *http.Request)
	UpdateOrder(writer http.ResponseWriter, request *http.Request)
	DeleteOrder(writer http.ResponseWriter, request *http.Request)
}
