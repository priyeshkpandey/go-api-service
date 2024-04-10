package service

import (
	"encoding/json"
	"io"
	"net/http"

	"api.service.go/go-api-service/common"
	"api.service.go/go-api-service/model"
	"api.service.go/go-api-service/repository"
	"github.com/gorilla/mux"
)

const dbFilePath string = "./test.db"

type OrderService struct {
	OrderRep *repository.OrderRepo
}

func NewOrderService() *OrderService {
	service := &OrderService{}
	service.OrderRep = repository.NewOrderRepo(dbFilePath)
	return service
}

func (service *OrderService) GetAllOrders(writer http.ResponseWriter, reader *http.Request) {
	if orders := service.OrderRep.GetAll(); orders != nil {
		buildJsonResponse(writer, http.StatusOK, orders)
	} else {
		buildErrorJsonResponse(writer, http.StatusNotFound, "No orders found")
	}

}

func (service *OrderService) GetOrder(writer http.ResponseWriter, reader *http.Request) {
	pathVariables := mux.Vars(reader)
	orderId := pathVariables["id"]
	if order := service.OrderRep.GetOrder(orderId); order != nil {
		buildJsonResponse(writer, http.StatusOK, order)
	} else {
		buildErrorJsonResponse(writer, http.StatusNotFound, "No order found")
	}
}

func (service *OrderService) CreateOrder(writer http.ResponseWriter, reader *http.Request) {
	reqBody, err := io.ReadAll(reader.Body)
	if !common.HasError(err) {
		var order *model.Order
		err = json.Unmarshal(reqBody, order)
		if !common.HasError(err) {
			if orderResponse := service.OrderRep.CreateOrder(order); orderResponse != nil {
				buildJsonResponse(writer, http.StatusCreated, orderResponse)
			} else {
				buildErrorJsonResponse(writer, http.StatusBadRequest, "Encountered error in creating order")
			}
		} else {
			buildErrorJsonResponse(writer, http.StatusBadRequest, "Encountered error in parsing request body")
		}

	} else {
		buildErrorJsonResponse(writer, http.StatusInternalServerError, "Error reading request body.")
	}
}

func (service *OrderService) UpdateOrder(writer http.ResponseWriter, reader *http.Request) {
	reqBody, err := io.ReadAll(reader.Body)
	if !common.HasError(err) {
		var order *model.Order
		json.Unmarshal(reqBody, order)
		if orderResponse := service.OrderRep.UpdateOrder(order); orderResponse != nil {
			buildJsonResponse(writer, http.StatusCreated, orderResponse)
		} else {
			buildErrorJsonResponse(writer, http.StatusBadRequest, "Encountered error in updating order")
		}
	} else {
		buildErrorJsonResponse(writer, http.StatusInternalServerError, "Error reading request body.")
	}
}

func (service *OrderService) DeleteOrder(writer http.ResponseWriter, reader *http.Request) {
	reqBody, err := io.ReadAll(reader.Body)
	if !common.HasError(err) {
		var order *model.Order
		json.Unmarshal(reqBody, order)
		if orderResponse := service.OrderRep.DeleteOrder(order); orderResponse {
			buildJsonResponse(writer, http.StatusCreated, orderResponse)
		} else {
			buildErrorJsonResponse(writer, http.StatusBadRequest, "Encountered error in deleting order")
		}
	} else {
		buildErrorJsonResponse(writer, http.StatusInternalServerError, "Error reading request body.")
	}
}

func (service *OrderService) Close() {
	service.OrderRep.Close()
}

func buildErrorJsonResponse(writer http.ResponseWriter, code int, message string) {
	buildJsonResponse(writer, code, map[string]string{"error": message})
}

func buildJsonResponse(writer http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if !common.HasError(err) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(code)
		writer.Write(response)
	} else {
		writer.Write([]byte(`Error creating json response`))
	}

}
