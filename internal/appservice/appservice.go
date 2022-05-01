package appservice

import (
	"context"
	"errors"
	"fmt"
	"layers/internal/app"
	"layers/internal/storage"
)

type App struct {
	storage         storage.Facade
	itemTransformer *itemTransformer
}

var _ app.Facade = (*App)(nil) // verify interface compliance

func NewApp(storage storage.Facade, itemTransformer *itemTransformer) app.Facade {
	return &App{
		storage:         storage,
		itemTransformer: itemTransformer,
	}
}

func (s *App) AddItem(ctx context.Context, item app.Item) error {
	return s.storage.AddItem(ctx, s.itemTransformer.ToStorageItem(item))
}

func (s *App) GetItem(ctx context.Context, key string) (app.Item, error) {
	storageItem, err := s.storage.GetItemByKey(ctx, key)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return app.Item{}, nil
		}
		return app.Item{}, fmt.Errorf("get item by key %s: %v", key, err)
	}

	return s.itemTransformer.FromStorageItem(storageItem), nil
}
