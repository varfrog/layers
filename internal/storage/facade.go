package storage

import (
	"context"

	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

type Item struct {
	Key string
	Val string
}

type Facade interface {
	// GetItemByKey returns an Item from storage. If not found, returns an error ErrNotFound.
	GetItemByKey(ctx context.Context, key string) (Item, error)

	// AddItem adds an item to storage, at key equal to the item's Key field. If an item exists at this key,
	// it will be overridden.
	AddItem(ctx context.Context, item Item) error
}
