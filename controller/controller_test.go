package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"redis/model"
	servicemocks "redis/service/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_ControllerGetHumans(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("GetHumans should return humans", func(t *testing.T) {
		ctx, rec := getRequestForHumans()

		mockService := servicemocks.NewMockService(mockController)
		mockService.EXPECT().GetHumans(gomock.Any()).Return(&model.Humans{
			model.Human{
				ID:   primitive.ObjectID{},
				Name: "John Doe",
				Age:  35,
			},
		}, nil).Times(1)

		serverController := NewController(mockService)
		err := serverController.GetHumans(ctx)
		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("GetHumans should return error", func(t *testing.T) {
		ctx, rec := getRequestForHumans()

		mockService := servicemocks.NewMockService(mockController)
		mockService.EXPECT().GetHumans(gomock.Any()).Return(
			&model.Humans{},
			errors.New("error"),
		).Times(1)

		serverController := NewController(mockService)
		err := serverController.GetHumans(ctx)
		require.Nil(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}

func Test_ControllerGetHuman(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("GetHuman should return human", func(t *testing.T) {
		var err error
		ID := strings.Repeat("1", 24)
		objectID, err := primitive.ObjectIDFromHex(ID)
		require.Nil(t, err)

		expectedHuman := model.Human{
			ID:   objectID,
			Name: "John Doe",
			Age:  35,
		}

		mockService := servicemocks.NewMockService(mockController)
		mockService.EXPECT().GetHuman(
			gomock.Any(),
			gomock.Eq("123456"),
		).Return(&expectedHuman, nil).Times(1)

		ctx, rec := getRequestForHuman()
		serverController := NewController(mockService)
		err = serverController.GetHuman(ctx)

		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var actualHuman model.Human
		err = json.Unmarshal(rec.Body.Bytes(), &actualHuman)

		require.Nil(t, err)
		assert.Equal(t, expectedHuman, actualHuman)
	})

	t.Run("GetHuman should return error", func(t *testing.T) {
		var err error
		mockService := servicemocks.NewMockService(mockController)
		mockService.EXPECT().GetHuman(
			gomock.Any(),
			"123456",
		).Return(&model.Human{}, errors.New("error")).Times(1)

		ctx, _ := getRequestForHuman()
		serverController := NewController(mockService)
		err = serverController.GetHuman(ctx)

		require.Error(t, err)
	})
}

func getRequestForHumans() (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/humans", nil)
	req.Header.Set("Accept", "application/json")

	rec := httptest.NewRecorder()

	ctx := echo.New().NewContext(req, rec)

	return ctx, rec
}

func getRequestForHuman() (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/humans/:id", nil)
	req.Header.Set("Accept", "application/json")

	rec := httptest.NewRecorder()

	ctx := echo.New().NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("123456")

	return ctx, rec
}
