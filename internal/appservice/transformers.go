package appservice

import (
	"layers/internal/app"
	"layers/internal/storage"
)

type itemTransformer struct{}

func NewItemTransformer() *itemTransformer {
	return &itemTransformer{}
}

func (s *itemTransformer) ToStorageItem(item app.Item) storage.Item {
	return storage.Item{
		Key: item.Key,
		Val: item.Val,
	}
}

func (s *itemTransformer) FromStorageItem(storageItem storage.Item) app.Item {
	return app.Item{
		Key: storageItem.Key,
		Val: storageItem.Val,
	}
}
