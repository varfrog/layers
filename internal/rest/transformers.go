package rest

import (
	"layers/internal/app"
)

type ItemTransformer struct{}

func NewItemTransformer() *ItemTransformer {
	return &ItemTransformer{}
}

func (t *ItemTransformer) ToAppItem(request AddItemRequest) app.Item {
	return app.Item{
		Key: request.Key,
		Val: request.Val,
	}
}
