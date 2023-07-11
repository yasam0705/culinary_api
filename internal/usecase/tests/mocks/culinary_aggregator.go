// Code generated by MockGen. DO NOT EDIT.
// Source: culinary_aggregator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "github/culinary_api/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCulinaryAggregator is a mock of CulinaryAggregator interface.
type MockCulinaryAggregator struct {
	ctrl     *gomock.Controller
	recorder *MockCulinaryAggregatorMockRecorder
}

// MockCulinaryAggregatorMockRecorder is the mock recorder for MockCulinaryAggregator.
type MockCulinaryAggregatorMockRecorder struct {
	mock *MockCulinaryAggregator
}

// NewMockCulinaryAggregator creates a new mock instance.
func NewMockCulinaryAggregator(ctrl *gomock.Controller) *MockCulinaryAggregator {
	mock := &MockCulinaryAggregator{ctrl: ctrl}
	mock.recorder = &MockCulinaryAggregatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCulinaryAggregator) EXPECT() *MockCulinaryAggregatorMockRecorder {
	return m.recorder
}

// AddRating mocks base method.
func (m *MockCulinaryAggregator) AddRating(ctx context.Context, userId, recipeId string, rating int8) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRating", ctx, userId, recipeId, rating)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRating indicates an expected call of AddRating.
func (mr *MockCulinaryAggregatorMockRecorder) AddRating(ctx, userId, recipeId, rating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRating", reflect.TypeOf((*MockCulinaryAggregator)(nil).AddRating), ctx, userId, recipeId, rating)
}

// CreateCookingStep mocks base method.
func (m_2 *MockCulinaryAggregator) CreateCookingStep(ctx context.Context, m *entity.CookingSteps) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateCookingStep", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCookingStep indicates an expected call of CreateCookingStep.
func (mr *MockCulinaryAggregatorMockRecorder) CreateCookingStep(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCookingStep", reflect.TypeOf((*MockCulinaryAggregator)(nil).CreateCookingStep), ctx, m)
}

// CreateIngredient mocks base method.
func (m_2 *MockCulinaryAggregator) CreateIngredient(ctx context.Context, recipeId string, m *entity.Ingredients) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateIngredient", ctx, recipeId, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIngredient indicates an expected call of CreateIngredient.
func (mr *MockCulinaryAggregatorMockRecorder) CreateIngredient(ctx, recipeId, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIngredient", reflect.TypeOf((*MockCulinaryAggregator)(nil).CreateIngredient), ctx, recipeId, m)
}

// CreateRecipe mocks base method.
func (m_2 *MockCulinaryAggregator) CreateRecipe(ctx context.Context, m *entity.CulinaryAggregator) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateRecipe", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRecipe indicates an expected call of CreateRecipe.
func (mr *MockCulinaryAggregatorMockRecorder) CreateRecipe(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecipe", reflect.TypeOf((*MockCulinaryAggregator)(nil).CreateRecipe), ctx, m)
}

// DeleteCookingStep mocks base method.
func (m *MockCulinaryAggregator) DeleteCookingStep(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCookingStep", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCookingStep indicates an expected call of DeleteCookingStep.
func (mr *MockCulinaryAggregatorMockRecorder) DeleteCookingStep(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCookingStep", reflect.TypeOf((*MockCulinaryAggregator)(nil).DeleteCookingStep), ctx, id)
}

// DeleteIngredient mocks base method.
func (m *MockCulinaryAggregator) DeleteIngredient(ctx context.Context, recipeId, ingridientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIngredient", ctx, recipeId, ingridientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIngredient indicates an expected call of DeleteIngredient.
func (mr *MockCulinaryAggregatorMockRecorder) DeleteIngredient(ctx, recipeId, ingridientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIngredient", reflect.TypeOf((*MockCulinaryAggregator)(nil).DeleteIngredient), ctx, recipeId, ingridientId)
}

// DeleteRecipe mocks base method.
func (m *MockCulinaryAggregator) DeleteRecipe(ctx context.Context, filter map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecipe", ctx, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecipe indicates an expected call of DeleteRecipe.
func (mr *MockCulinaryAggregatorMockRecorder) DeleteRecipe(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecipe", reflect.TypeOf((*MockCulinaryAggregator)(nil).DeleteRecipe), ctx, filter)
}

// GetRecipe mocks base method.
func (m *MockCulinaryAggregator) GetRecipe(ctx context.Context, filter map[string]string) (*entity.CulinaryAggregator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecipe", ctx, filter)
	ret0, _ := ret[0].(*entity.CulinaryAggregator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipe indicates an expected call of GetRecipe.
func (mr *MockCulinaryAggregatorMockRecorder) GetRecipe(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipe", reflect.TypeOf((*MockCulinaryAggregator)(nil).GetRecipe), ctx, filter)
}

// UpdateCookingStep mocks base method.
func (m_2 *MockCulinaryAggregator) UpdateCookingStep(ctx context.Context, m *entity.CookingSteps) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateCookingStep", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCookingStep indicates an expected call of UpdateCookingStep.
func (mr *MockCulinaryAggregatorMockRecorder) UpdateCookingStep(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCookingStep", reflect.TypeOf((*MockCulinaryAggregator)(nil).UpdateCookingStep), ctx, m)
}

// UpdateIngredient mocks base method.
func (m_2 *MockCulinaryAggregator) UpdateIngredient(ctx context.Context, recipeId string, m *entity.Ingredients) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateIngredient", ctx, recipeId, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIngredient indicates an expected call of UpdateIngredient.
func (mr *MockCulinaryAggregatorMockRecorder) UpdateIngredient(ctx, recipeId, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIngredient", reflect.TypeOf((*MockCulinaryAggregator)(nil).UpdateIngredient), ctx, recipeId, m)
}

// UpdateRecipe mocks base method.
func (m_2 *MockCulinaryAggregator) UpdateRecipe(ctx context.Context, m *entity.CulinaryAggregator) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateRecipe", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRecipe indicates an expected call of UpdateRecipe.
func (mr *MockCulinaryAggregatorMockRecorder) UpdateRecipe(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecipe", reflect.TypeOf((*MockCulinaryAggregator)(nil).UpdateRecipe), ctx, m)
}

// MockCulinaryAggregatorRepo is a mock of CulinaryAggregatorRepo interface.
type MockCulinaryAggregatorRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCulinaryAggregatorRepoMockRecorder
}

// MockCulinaryAggregatorRepoMockRecorder is the mock recorder for MockCulinaryAggregatorRepo.
type MockCulinaryAggregatorRepoMockRecorder struct {
	mock *MockCulinaryAggregatorRepo
}

// NewMockCulinaryAggregatorRepo creates a new mock instance.
func NewMockCulinaryAggregatorRepo(ctrl *gomock.Controller) *MockCulinaryAggregatorRepo {
	mock := &MockCulinaryAggregatorRepo{ctrl: ctrl}
	mock.recorder = &MockCulinaryAggregatorRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCulinaryAggregatorRepo) EXPECT() *MockCulinaryAggregatorRepoMockRecorder {
	return m.recorder
}

// Ingridients mocks base method.
func (m *MockCulinaryAggregatorRepo) Ingridients(ctx context.Context, filters map[string]string) ([]*entity.Ingredients, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ingridients", ctx, filters)
	ret0, _ := ret[0].([]*entity.Ingredients)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ingridients indicates an expected call of Ingridients.
func (mr *MockCulinaryAggregatorRepoMockRecorder) Ingridients(ctx, filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ingridients", reflect.TypeOf((*MockCulinaryAggregatorRepo)(nil).Ingridients), ctx, filters)
}