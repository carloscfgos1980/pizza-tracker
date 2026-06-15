package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine, h *handler) {
	router.GET("/", h.ServeNewOrderForm)
	router.POST("/new-order", h.HandlerNewOrderPost)
	router.GET("/customer/:id", h.ServeCustomerOrder)

	router.Static("/static", "/templates/static")
}
