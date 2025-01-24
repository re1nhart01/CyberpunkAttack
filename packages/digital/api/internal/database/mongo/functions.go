package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (db *MongoDB) CreateCollection(ctx context.Context, name string, size int64, capped bool) error {
	opts := options.CreateCollection().SetCapped(capped)

	if capped {
		opts.SetSizeInBytes(size)
	}

	err := db.Get().CreateCollection(ctx, name, opts)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
