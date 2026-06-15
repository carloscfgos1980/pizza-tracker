package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{"Order Placed", "Preparing", "Quality Check", "Ready"}
	PizzaSizes    = []string{"Small", "Medium", "Large", "X Large"}
	PizzaTypes    = []string{"Margherita", "Pepperoni", "BBQ Chicken", "Veggie", "Hawaiian"}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID              string      `gorm:"primaryKey;size:14" json:"id"`
	Status          string      `gorm:"not null" json:"status"`
	CustomerName    string      `gorm:"not null" json:"customer_name"`
	CustomerPhone   string      `gorm:"not null" json:"customer_phone"`
	CustomerAddress string      `gorm:"not null" json:"customer_address"`
	Items           []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt       time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID           string `gorm:"primaryKey;size:14" json:"id"`
	OrderID      string `gorm:"index;size:14;not null" json:"order_id"`
	Size         string `gorm:"not null" json:"size"`
	Pizza        string `gorm:"not null" json:"pizza_id"`
	Instructions string `json:"instructions"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}

func (om *OrderModel) CreateOrder(order *Order) error {
	return om.DB.Create(order).Error
}

func (om *OrderModel) GetOrder(id string) (*Order, error) {
	var order Order
	if err := om.DB.Preload("Items").First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
