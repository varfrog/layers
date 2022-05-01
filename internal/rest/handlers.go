package rest

import (
	"context"
	"layers/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RequestHandler struct {
	serviceFacade   app.Facade
	itemTransformer *ItemTransformer
	logger          *zap.Logger
}

func NewRequestHandler(
	serviceFacade app.Facade,
	itemTransformer *ItemTransformer,
	logger *zap.Logger,
) *RequestHandler {
	return &RequestHandler{
		serviceFacade:   serviceFacade,
		itemTransformer: itemTransformer,
		logger:          logger,
	}
}

func (h *RequestHandler) HandleAddItemRequest(c *gin.Context) {
	var request AddItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid format"})
		return
	}

	err := h.serviceFacade.AddItem(context.Background(), h.itemTransformer.ToAppItem(request))
	if err != nil {
		h.logger.Error("add item", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Something went wrong"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"description": "Added"})
}

func (h *RequestHandler) HandleGetItemRequest(c *gin.Context) {
	var request GetItemRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid request"})
		return
	}

	item, err := h.serviceFacade.GetItem(context.Background(), request.Key)
	if err != nil {
		h.logger.Error("get item", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Could not get an Item"})
		return
	}

	if item == (app.Item{}) {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, ItemResponse{
		Key: item.Key,
		Val: item.Val,
	})
}
