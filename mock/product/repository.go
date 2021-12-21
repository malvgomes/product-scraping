// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/product/repository.go

// Package product_mock is a generated GoMock package.
package product_mock

import (
	product "product-scraping/pkg/domain/product"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetProductData mocks base method.
func (m *MockRepository) GetProductData(url string) (*product.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductData", url)
	ret0, _ := ret[0].(*product.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductData indicates an expected call of GetProductData.
func (mr *MockRepositoryMockRecorder) GetProductData(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductData", reflect.TypeOf((*MockRepository)(nil).GetProductData), url)
}

// InsertProductData mocks base method.
func (m *MockRepository) InsertProductData(data *product.Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProductData", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertProductData indicates an expected call of InsertProductData.
func (mr *MockRepositoryMockRecorder) InsertProductData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProductData", reflect.TypeOf((*MockRepository)(nil).InsertProductData), data)
}
