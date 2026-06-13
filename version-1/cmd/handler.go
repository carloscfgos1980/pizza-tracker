package main

import (
	"github.com/carloscfgos1980/pizza-tracker/internal/models"
)

type Handler struct {
	orders *models.OrderModel
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		orders: &dbModel.Order,
	}
}
