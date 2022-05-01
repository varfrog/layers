package rest

import "github.com/gin-gonic/gin"

func CreateRoutes(router *gin.Engine, requestHandler *RequestHandler) {
	router.POST("/items", requestHandler.HandleAddItemRequest)
	router.GET("/items/:key", requestHandler.HandleGetItemRequest)
}
