package mocks

import (
	"context"
	"database/sql"

	"github.com/imhasandl/search-service/internal/database"
	"github.com/stretchr/testify/mock"
)

// MockQueries provides a mock implementation of the DatabaseQuerier interface for testing.
// It uses the testify mock package to allow stubbing return values and tracking calls.
type MockQueries struct {
	mock.Mock
}

// NewMockQueries creates and returns a new instance of MockQueries for testing.
func NewMockQueries() *MockQueries {
	return &MockQueries{}
}

// SearchUsers mocks the SearchUsers method of the database interface.
// It returns users matching the provided query string.
func (m *MockQueries) SearchUsers(ctx context.Context, query sql.NullString) ([]database.User, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.User), args.Error(1)
}

// SearchUsersByDate mocks the SearchUsersByDate method of the database interface.
// It returns users matching the provided query string ordered by created_at timestamp.
func (m *MockQueries) SearchUsersByDate(ctx context.Context, query sql.NullString) ([]database.User, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.User), args.Error(1)
}

// SearchPosts mocks the SearchPosts method of the database interface.
// It returns posts that match the provided query string.
func (m *MockQueries) SearchPosts(ctx context.Context, query sql.NullString) ([]database.Post, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Post), args.Error(1)
}

// SearchPostsByDate mocks the SearchPostsByDate method of the database interface.
// It returns posts matching the provided query string ordered by created_at timestamp.
func (m *MockQueries) SearchPostsByDate(ctx context.Context, query sql.NullString) ([]database.Post, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Post), args.Error(1)
}

// SearchReports mocks the SearchReports method of the database interface.
// It returns reports that match the provided query string.
func (m *MockQueries) SearchReports(ctx context.Context, query sql.NullString) ([]database.Report, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Report), args.Error(1)
}

// SearchReportsByDate mocks the SearchReportsByDate method of the database interface.
// It returns reports matching the provided query string ordered by reported_at timestamp.
func (m *MockQueries) SearchReportsByDate(ctx context.Context, query sql.NullString) ([]database.Report, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]database.Report), args.Error(1)
}
