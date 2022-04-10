package cache

import (
	"context"
	"redis/model"
)

type Cache interface {
	Set(ctx context.Context, key string, value *model.Human) error
	Get(ctx context.Context, key string) (*model.Human, error)
}
