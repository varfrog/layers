package app

import "context"

type Item struct {
	Key string
	Val string
}

type Facade interface {
	// AddItem adds an item to the underlying storage, at key equal to the item's Key field.
	// Overrides the value if an item exists at this key.
	AddItem(ctx context.Context, item Item) error

	// GetItem returns an Item from the underlying storage whose key equals the given key.
	// If no such item exists, it returns an empty Item.
	GetItem(ctx context.Context, key string) (Item, error)
}
