package repository

import (
	"database/sql"
	"fmt"
	"log"

	"api.service.go/go-api-service/common"
	"api.service.go/go-api-service/database"
	"api.service.go/go-api-service/model"
)

type OrderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(dbFilePath string) *OrderRepo {
	orderRepo := &OrderRepo{}
	orderRepo.DB = database.OpenDB(dbFilePath)
	return orderRepo
}

func (orderRepo *OrderRepo) GetAll() []*model.Order {
	rows, err := orderRepo.DB.Query("SELECT id, product_ids, order_amount FROM product_order")
	var orders []*model.Order
	if !common.HasError(err) {
		defer rows.Close()
		for rows.NextResultSet() {
			rows.Next()
			order := model.NewOrder()
			err = rows.Scan(order.ID, order.ProductIDs, order.OrderAmount)
			if !common.HasError(err) {
				orders = append(orders, order)
			}

		}
	}
	return orders
}

func (orderRepo *OrderRepo) GetOrder(orderId string) *model.Order {
	rows := orderRepo.DB.QueryRow("SELECT id, product_ids, order_amount FROM product_order where id = ?", orderId)
	order := model.NewOrder()
	err := rows.Scan(order.ID, order.ProductIDs, order.OrderAmount)
	if !common.HasError(err) {
		return order
	}
	return nil
}

func (orderRepo *OrderRepo) CreateOrder(order *model.Order) *model.Order {
	log.Printf("Received order: %v", *order)
	result, err := orderRepo.DB.Exec("INSERT INTO product_order(id, product_ids, order_amount) VALUES(?, ?, ?)", order.ID, order.ProductIDs, order.OrderAmount)
	if !common.HasError(err) {
		numOfRows, err := result.RowsAffected()
		log.Printf("Number of rows inserted: %d", numOfRows)
		if !common.HasError(err) {
			return orderRepo.GetOrder(fmt.Sprint(order.ID))
		}
	}
	return nil
}

func (orderRepo *OrderRepo) UpdateOrder(order *model.Order) *model.Order {
	result, err := orderRepo.DB.Exec("UPDATE product_order SET product_ids = ?, order_amount = ? WHERE id = ?", order.ProductIDs, order.OrderAmount, order.ID)
	if !common.HasError(err) {
		numOfRows, err := result.RowsAffected()
		log.Printf("Number of rows updated: %d", numOfRows)
		if !common.HasError(err) {
			return orderRepo.GetOrder(fmt.Sprint(order.ID))
		}
	}
	return nil
}

func (orderRepo *OrderRepo) DeleteOrder(order *model.Order) bool {
	result, err := orderRepo.DB.Exec("DELETE FROM product_order WHERE id = ?", order.ID)
	if !common.HasError(err) {
		numOfRows, err := result.RowsAffected()
		log.Printf("Number of rows deleted: %d", numOfRows)
		if !common.HasError(err) {
			return true
		}
	}
	return false
}

func (orderRepo *OrderRepo) Close() {
	orderRepo.DB.Close()
}
