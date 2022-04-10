package service

import (
	"context"
	"errors"
	cachemocks "redis/cache/mocks"
	"redis/model"
	repositorymocks "redis/repository/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_ServiceGetHumans(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRepository := repositorymocks.NewMockRepository(mockController)

	expectedHumans := model.Humans{
		model.Human{
			ID:   primitive.ObjectID{},
			Name: "John Doe",
			Age:  35,
		},
		model.Human{
			ID:   primitive.ObjectID{},
			Name: "Jane Doe",
			Age:  35,
		},
	}

	mockRepository.EXPECT().GetHumans(gomock.Any()).Return(&expectedHumans, nil).Times(1)

	serverService := NewService(mockRepository, nil)

	actualHumans, err := serverService.GetHumans(context.Background())
	require.Nil(t, err)
	assert.Equal(t, expectedHumans, *actualHumans)
}

func Test_ServiceGetHuman(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("GetHuman should return human from cache", func(t *testing.T) {
		var err error
		ID := strings.Repeat("1", 24)
		objectID, err := primitive.ObjectIDFromHex(ID)
		require.Nil(t, err)

		expectedHuman := model.Human{
			ID:   objectID,
			Name: "John Doe",
			Age:  35,
		}

		mockCache := cachemocks.NewMockCache(mockController)

		mockCache.EXPECT().Get(gomock.Any(), gomock.Eq("123456")).Return(
			&expectedHuman,
			nil,
		).Times(1)

		serverService := NewService(nil, mockCache)

		actualHuman, err := serverService.GetHuman(context.Background(), "123456")
		require.Nil(t, err)
		assert.Equal(t, expectedHuman, *actualHuman)
	})

	t.Run("GetHuman should return human from repository", func(t *testing.T) {
		var err error

		ID := strings.Repeat("1", 24)
		objectID, err := primitive.ObjectIDFromHex(ID)
		require.Nil(t, err)

		expectedHuman := model.Human{
			ID:   objectID,
			Name: "John Doe",
			Age:  35,
		}

		mockRepository := repositorymocks.NewMockRepository(mockController)
		mockCache := cachemocks.NewMockCache(mockController)

		mockRepository.EXPECT().GetHuman(gomock.Any(), gomock.Eq("123456")).Return(
			&expectedHuman,
			nil,
		).Times(1)
		mockCache.EXPECT().Get(gomock.Any(), gomock.Eq("123456")).Return(
			nil,
			errors.New("error"),
		).Times(1)
		mockCache.EXPECT().Set(
			gomock.Any(),
			gomock.Eq("123456"),
			gomock.Eq(&expectedHuman),
		).Return(nil).Times(1)

		serverService := NewService(mockRepository, mockCache)

		actualHuman, err := serverService.GetHuman(context.Background(), "123456")
		require.Nil(t, err)
		assert.Equal(t, expectedHuman, *actualHuman)
	})
}
