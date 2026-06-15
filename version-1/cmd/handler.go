package main

import (
	"github.com/carloscfgos1980/pizza-tracker/internal/models"
)

type handler struct {
	orders *models.OrderModel
}

func NewHandler(dbModel *models.DBModel) *handler {
	return &handler{
		orders: &dbModel.Order,
	}
}
