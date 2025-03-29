package mocks

import (
	"context"
	"database/sql"

	"github.com/imhasandl/search-service/internal/database"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
	mock.Mock
}

func NewMockQueries() *MockQueries {
	return &MockQueries{}
}

func (m *MockQueries) SearchUsers(ctx context.Context, query sql.NullString) ([]database.User, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.User), args.Error(1)
}

func (m *MockQueries) SearchUsersByDate(ctx context.Context, query sql.NullString) ([]database.User, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.User), args.Error(1)
}

func (m *MockQueries) SearchPosts(ctx context.Context, query sql.NullString) ([]database.Post, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Post), args.Error(1)
}

func (m *MockQueries) SearchPostsByDate(ctx context.Context, query sql.NullString) ([]database.Post, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Post), args.Error(1)
}

func (m *MockQueries) SearchReports(ctx context.Context, query sql.NullString) ([]database.Report, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Report), args.Error(1)
}

func (m *MockQueries) SearchReportsByDate(ctx context.Context, query sql.NullString) ([]database.Report, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Report), args.Error(1)
}