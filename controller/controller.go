package controller

import (
	"log"
	"net/http"

	"api.service.go/go-api-service/service"
	"github.com/gorilla/mux"
)

func getAllOrders(writer http.ResponseWriter, reader *http.Request) {
	service := service.NewOrderService()
	defer service.Close()
	service.GetAllOrders(writer, reader)
}

func getOrder(writer http.ResponseWriter, reader *http.Request) {
	service := service.NewOrderService()
	defer service.Close()
	service.GetOrderById(writer, reader)
}

func createOrder(writer http.ResponseWriter, reader *http.Request) {
	service := service.NewOrderService()
	defer service.Close()
	service.CreateOrder(writer, reader)
}

func updateOrder(writer http.ResponseWriter, reader *http.Request) {
	service := service.NewOrderService()
	defer service.Close()
	service.UpdateOrder(writer, reader)
}

func deleteOrder(writer http.ResponseWriter, reader *http.Request) {
	service := service.NewOrderService()
	defer service.Close()
	service.DeleteOrder(writer, reader)
}

type Controller struct {
	Router *mux.Router
	Port   string
}

func NewController(port string) *Controller {
	controller := &Controller{}
	controller.Router = mux.NewRouter()
	controller.Port = port
	return controller
}

func (controller *Controller) InitController() {
	controller.Router.HandleFunc("/orders", getAllOrders).Methods("GET")
	controller.Router.HandleFunc("/orders/{id}", getOrder).Methods("GET")
	controller.Router.HandleFunc("/orders", createOrder).Methods("POST")
	controller.Router.HandleFunc("/orders", updateOrder).Methods("PUT")
	controller.Router.HandleFunc("/orders", deleteOrder).Methods("DELETE")

	http.Handle("/", controller.Router)
	log.Printf("Server started and listening on localhost:%s", controller.Port)
	log.Fatal(http.ListenAndServe(":"+controller.Port, nil))
}
