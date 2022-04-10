package service

import (
	"context"
	"redis/cache"
	"redis/model"
	"redis/repository"
)

type Service interface {
	GetHumans(ctx context.Context) (*model.Humans, error)
	GetHuman(ctx context.Context, id string) (*model.Human, error)
}

type service struct {
	repository repository.Repository
	redisCache cache.Cache
}

func NewService(repository repository.Repository, redisCache cache.Cache) Service {
	return &service{
		repository: repository,
		redisCache: redisCache,
	}
}

func (s *service) GetHumans(ctx context.Context) (*model.Humans, error) {
	return s.repository.GetHumans(ctx)
}

func (s *service) GetHuman(ctx context.Context, id string) (*model.Human, error) {
	var err error
	human, err := s.redisCache.Get(ctx, id)
	if human == nil || err != nil {
		human, err = s.repository.GetHuman(ctx, id)
		if err != nil {
			return human, err
		}
		err = s.redisCache.Set(ctx, id, human)
		if err != nil {
			return human, err
		}
	}
	return human, err
}
