package handlers

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	productmock "product-scraping/mock/product"
	"product-scraping/pkg/domain/product"
	"strings"
	"testing"
)

func TestProductHandler_ParseProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		handler, mock, finish := getMockedHandler(t)
		defer finish()

		mock.EXPECT().ParseProductURL("https://url.com").
			Return(&product.Entity{
				Title:       "title",
				ImageURL:    "imageurl",
				Price:       200,
				Description: "desc",
				URL:         "https://url.com",
			}, nil)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "/product", strings.NewReader(`{"url": "https://url.com"}`))
		assert.Nil(t, err)

		handler.ParseProduct(w, r)

		result := w.Result()
		defer result.Body.Close()

		body, err := io.ReadAll(result.Body)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, `{"title":"title","imageURL":"imageurl","price":200,"description":"desc","url":"https://url.com"}`, string(body))
	})

	t.Run("Failure - Invalid request payload", func(t *testing.T) {
		handler, _, finish := getMockedHandler(t)
		defer finish()

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "/product", strings.NewReader(`{"url": ""}`))
		assert.Nil(t, err)

		handler.ParseProduct(w, r)

		result := w.Result()
		defer result.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	})

	t.Run("Failure - Parsing URL", func(t *testing.T) {
		handler, mock, finish := getMockedHandler(t)
		defer finish()

		mock.EXPECT().ParseProductURL("https://url.com").
			Return(nil, errors.New("failure"))

		w := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodPost, "/product", strings.NewReader(`{"url": "https://url.com"}`))
		assert.Nil(t, err)

		handler.ParseProduct(w, r)

		result := w.Result()
		defer result.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, result.StatusCode)
	})
}

func getMockedHandler(t *testing.T) (*ProductHandler, *productmock.MockUseCase, func()) {
	t.Helper()

	ctrl := gomock.NewController(t)

	finish := func() {
		ctrl.Finish()
	}

	usecaseMock := productmock.NewMockUseCase(ctrl)

	handler := NewProductHandler(usecaseMock)

	return handler, usecaseMock, finish
}
