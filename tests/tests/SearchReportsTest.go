package tests

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/imhasandl/search-service/cmd/server"
	"github.com/imhasandl/search-service/internal/database"
	pb "github.com/imhasandl/search-service/protos"
	"github.com/imhasandl/search-service/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSearchReports(t *testing.T) {
	// Setup
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")
	
	t.Run("successful search", func(t *testing.T) {
		 // Prepare test data
		 reportID := uuid.New()
		 userID := uuid.New()
		 testTime := time.Now()
		 query := "user"
		 nullQuery := sql.NullString{String: query, Valid: true}
		 
		 // Configure mock
		 mockDB.On("SearchReports", mock.Anything, nullQuery).Return([]database.Report{
			  {
					ID:         reportID,
					ReportedAt: testTime,
					ReportedBy: userID,
					Reason:     "Inappropriate content",
			  },
		 }, nil).Once()
		 
		 // Execute the method
		 resp, err := testServer.SearchReports(context.Background(), &pb.SearchReportsRequest{
			  Query: query,
		 })
		 
		 // Verify
		 assert.NoError(t, err)
		 assert.NotNil(t, resp)
		 assert.Equal(t, 1, len(resp.Report))
		 assert.Equal(t, reportID.String(), resp.Report[0].Id)
		 assert.Equal(t, userID.String(), resp.Report[0].ReportedBy)
		 assert.Equal(t, "Inappropriate content", resp.Report[0].Reason)
		 
		 // Verify mock expectations
		 mockDB.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		 // Configure mock to return error
		 query := "error"
		 nullQuery := sql.NullString{String: query, Valid: true}
		 
		 mockDB.On("SearchReports", mock.Anything, nullQuery).Return(
			  []database.Report{}, errors.New("database error"),
		 ).Once()
		 
		 // Execute the method
		 resp, err := testServer.SearchReports(context.Background(), &pb.SearchReportsRequest{
			  Query: query,
		 })
		 
		 // Verify error is returned correctly
		 assert.Error(t, err)
		 assert.Nil(t, resp)
		 
		 // Check if error is correctly formatted as gRPC error
		 statusErr, ok := status.FromError(err)
		 assert.True(t, ok)
		 assert.Equal(t, codes.Internal, statusErr.Code())
		 assert.Contains(t, statusErr.Message(), "can't get report")
		 
		 // Verify mock expectations
		 mockDB.AssertExpectations(t)
	})
}

func TestSearchReportsByDate(t *testing.T) {
	
}