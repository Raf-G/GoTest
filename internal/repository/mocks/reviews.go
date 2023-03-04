// Code generated by MockGen. DO NOT EDIT.
// Source: reviews.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	"example.com/m/v2/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReviewsStorage is a mock of ReviewsStorage interface.
type MockReviewsStorage struct {
	ctrl     *gomock.Controller
	recorder *MockReviewsStorageMockRecorder
}

// MockReviewsStorageMockRecorder is the mock recorder for MockReviewsStorage.
type MockReviewsStorageMockRecorder struct {
	mock *MockReviewsStorage
}

// NewMockReviewsStorage creates a new mock instance.
func NewMockReviewsStorage(ctrl *gomock.Controller) *MockReviewsStorage {
	mock := &MockReviewsStorage{ctrl: ctrl}
	mock.recorder = &MockReviewsStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewsStorage) EXPECT() *MockReviewsStorageMockRecorder {
	return m.recorder
}

// AddReview mocks base method.
func (m *MockReviewsStorage) AddReview(arg0 domain.Review) (*domain.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReview", arg0)
	ret0, _ := ret[0].(*domain.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddReview indicates an expected call of AddReview.
func (mr *MockReviewsStorageMockRecorder) AddReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReview", reflect.TypeOf((*MockReviewsStorage)(nil).AddReview), arg0)
}

// DeleteReview mocks base method.
func (m *MockReviewsStorage) DeleteReview(arg0 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockReviewsStorageMockRecorder) DeleteReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockReviewsStorage)(nil).DeleteReview), arg0)
}

// EditReview mocks base method.
func (m *MockReviewsStorage) EditReview(arg0 domain.Review) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditReview", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditReview indicates an expected call of EditReview.
func (mr *MockReviewsStorageMockRecorder) EditReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditReview", reflect.TypeOf((*MockReviewsStorage)(nil).EditReview), arg0)
}

// GetReview mocks base method.
func (m *MockReviewsStorage) GetReview(arg0 int) (domain.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", arg0)
	ret0, _ := ret[0].(domain.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockReviewsStorageMockRecorder) GetReview(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockReviewsStorage)(nil).GetReview), arg0)
}

// GetReviews mocks base method.
func (m *MockReviewsStorage) GetReviews() ([]domain.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviews")
	ret0, _ := ret[0].([]domain.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviews indicates an expected call of GetReviews.
func (mr *MockReviewsStorageMockRecorder) GetReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviews", reflect.TypeOf((*MockReviewsStorage)(nil).GetReviews))
}

// GetReviewsProduct mocks base method.
func (m *MockReviewsStorage) GetReviewsProduct(arg0 int) ([]domain.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsProduct", arg0)
	ret0, _ := ret[0].([]domain.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewsProduct indicates an expected call of GetReviewsProduct.
func (mr *MockReviewsStorageMockRecorder) GetReviewsProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsProduct", reflect.TypeOf((*MockReviewsStorage)(nil).GetReviewsProduct), arg0)
}
