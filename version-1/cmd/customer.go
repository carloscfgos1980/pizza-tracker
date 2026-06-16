package main

import (
	"log/slog"
	"net/http"

	"github.com/carloscfgos1980/pizza-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

type CustomerData struct {
	title    string
	Order    models.Order
	Statuses []string
}

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name         string   `form:"name" binding:"required,min=2,max=100"`
	Phone        string   `form:"phone" binding:"required,min=10,max=15"`
	Address      string   `form:"address" binding:"required,min=5,max=200"`
	Size         []string `form:"size" binding:"required,min=1,dive,pizza_valid_size"`
	PizzaType    []string `form:"pizza" binding:"required,min=1,dive,pizza_valid_type"`
	Instructions []string `form:"instructions" binding:"omitempty,max=200"`
}

func (h *handler) ServeNewOrderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		PizzaTypes: models.PizzaTypes,
		PizzaSizes: models.PizzaSizes,
	})
}

func (h *handler) HandlerNewOrderPost(c *gin.Context) {
	var form OrderRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItems := make([]models.OrderItem, len(form.PizzaType))
	for i := range form.PizzaType {
		orderItems[i] = models.OrderItem{
			Pizza:        form.PizzaType[i],
			Size:         form.Size[i],
			Instructions: form.Instructions[i],
		}
	}
	order := models.Order{
		Status:          models.OrderStatuses[0],
		CustomerName:    form.Name,
		CustomerPhone:   form.Phone,
		CustomerAddress: form.Address,
		Items:           orderItems,
	}
	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Somethig went wrong")
		return
	}
	slog.Info("Order created successfully", "order_id", order.ID, "customer", order.CustomerName)
	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)
}

func (h *handler) ServeCustomerOrder(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.String(http.StatusBadRequest, "Order ID is required")
		return
	}
	order, err := h.orders.GetOrder(orderID)
	if err != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}
	c.HTML(http.StatusOK, "customer.tmpl", CustomerData{
		title:    "Pizza Order Status - " + order.ID,
		Order:    *order,
		Statuses: models.OrderStatuses,
	})
}
