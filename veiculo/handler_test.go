package veiculo_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"rentcar/veiculo"
	"rentcar/veiculo/mocks"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestGet(t *testing.T) {
	mockstg := &mocks.Storage{}

	targetURL := "/veiculos"
	expectedList := []veiculo.Veiculo{
		veiculo.Veiculo{
			ID:     1,
			Nome:   "UP TSI",
			Marca:  "VW",
			Modelo: 2018,
			Ano:    2018,
		},
		veiculo.Veiculo{
			ID:     2,
			Nome:   "JETTA TSI",
			Marca:  "VW",
			Modelo: 2019,
			Ano:    2019,
		},
	}

	t.Run("GET - success", func(t *testing.T) {
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)
		mockstg.On("GetVeiculos").Return(expectedList, nil)
		handler := veiculo.NewVeiculo(mockstg)
		r.GET(targetURL, handler.Get)
		req := httptest.NewRequest(http.MethodGet, targetURL, nil)

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		b, err := ioutil.ReadAll(w.Body)
		assert.NoError(t, err)

		var actual []veiculo.Veiculo
		err = json.Unmarshal(b, &actual)
		assert.NoError(t, err)
		assert.Equal(t, expectedList, actual)
	})
}

func TestCreateOrUpdate(t *testing.T) {
	veicultoTest := veiculo.Veiculo{
		ID:     1,
		Nome:   "UP TSI",
		Marca:  "VW",
		Ano:    2018,
		Modelo: 2018,
	}

	stgmock := &mocks.Storage{}
	t.Run("POST - success", func(t *testing.T) {
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)
		stgmock.On("CreateVeiculo", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
		handler := veiculo.NewVeiculo(stgmock)
		r.POST("/veiculos", handler.Create)
		jv, err := json.Marshal(veicultoTest)

		if !assert.NoError(t, err) {
			t.FailNow()
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/veiculos", bytes.NewBuffer(jv))
		r.ServeHTTP(w, req)
		assert.Equal(t, 201, w.Code)
	})

	t.Run("PUT - Update veiculo", func(t *testing.T) {
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)
		stgmock.On("UpdateVeiculo", mock.Anything, mock.Anything).Return(nil)
		handler := veiculo.NewVeiculo(stgmock)
		r.PUT("/veiculos", handler.Update)
		jv, err := json.Marshal(veicultoTest)

		if !assert.NoError(t, err) {
			t.FailNow()
			panic(err)
		}

		req := httptest.NewRequest(http.MethodPut, "/veiculos", bytes.NewBuffer(jv))
		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}

func TestDelete(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	stgMock := &mocks.Storage{}
	stgMock.On("DeleteVeiculo", mock.Anything).Return(nil).Once()
	handler := veiculo.NewVeiculo(stgMock)
	r.DELETE("/veiculos/:id", handler.Delete)

	req := httptest.NewRequest(http.MethodDelete, "/veiculos/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
