package tests

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/imhasandl/search-service/cmd/server"
	"github.com/imhasandl/search-service/internal/mocks"
	"github.com/imhasandl/search-service/internal/database"
	pb "github.com/imhasandl/search-service/protos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSearchReports(t *testing.T) {
	// Test setup - common for all test cases
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")
	testTime := time.Now()

	// Define test cases
	testCases := []struct {
		name           string
		query          string
		mockSetup      func()
		expectedError  bool
		expectedCode   codes.Code
		expectedErrMsg string
		validateResp   func(t *testing.T, resp *pb.SearchReportsResponse)
	}{
		{
			name:  "successful search",
			query: "spam",
			mockSetup: func() {
				reportID := uuid.New()
				userID := uuid.New()
				nullQuery := sql.NullString{String: "spam", Valid: true}

				mockDB.On("SearchReports", mock.Anything, nullQuery).Return([]database.Report{
					{
						ID:         reportID,
						ReportedAt: testTime,
						ReportedBy: userID,
						Reason:     "This is spam content",
					},
				}, nil).Once()
			},
			expectedError: false,
			validateResp: func(t *testing.T, resp *pb.SearchReportsResponse) {
				assert.NotNil(t, resp)
				assert.Equal(t, 1, len(resp.Report))
				assert.Contains(t, resp.Report[0].Reason, "spam")
			},
		},
		{
			name:  "empty query",
			query: "",
			mockSetup: func() {
				nullQuery := sql.NullString{String: "", Valid: false}
				mockDB.On("SearchReports", mock.Anything, nullQuery).Return([]database.Report{}, nil).Once()
			},
			expectedError: false,
			validateResp: func(t *testing.T, resp *pb.SearchReportsResponse) {
				assert.NotNil(t, resp)
				assert.Equal(t, 0, len(resp.Report))
			},
		},
		{
			name:  "database error",
			query: "error",
			mockSetup: func() {
				nullQuery := sql.NullString{String: "error", Valid: true}
				mockDB.On("SearchReports", mock.Anything, nullQuery).Return(
					[]database.Report{}, errors.New("database error"),
				).Once()
			},
			expectedError:  true,
			expectedCode:   codes.Internal,
			expectedErrMsg: "can't get report",
			validateResp:   nil,
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock for this test case
			tc.mockSetup()

			// Execute the method
			resp, err := testServer.SearchReports(context.Background(), &pb.SearchReportsRequest{
				Query: tc.query,
			})

			// Validate results
			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, resp)

				// Check if error is correctly formatted as gRPC error
				statusErr, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tc.expectedCode, statusErr.Code())
				assert.Contains(t, statusErr.Message(), tc.expectedErrMsg)
			} else {
				assert.NoError(t, err)
				tc.validateResp(t, resp)
			}

			// Verify mock expectations
			mockDB.AssertExpectations(t)
		})
	}
}

func TestSearchReportsByDate(t *testing.T) {
	// Test setup
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")

	// Define test cases
	testCases := []struct {
		name         string
		query        string
		mockSetup    func() (uuid.UUID, uuid.UUID, time.Time, time.Time)
		validateResp func(t *testing.T, resp *pb.SearchReportsByDateResponse, reportID1, reportID2 uuid.UUID)
	}{
		{
			name:  "successful search with sorting",
			query: "inappropriate",
			mockSetup: func() (uuid.UUID, uuid.UUID, time.Time, time.Time) {
				reportID1 := uuid.New()
				reportID2 := uuid.New()
				userID := uuid.New()
				testTime1 := time.Now().Add(-24 * time.Hour) // older
				testTime2 := time.Now()                      // newer

				nullQuery := sql.NullString{String: "inappropriate", Valid: true}
				mockDB.On("SearchReportsByDate", mock.Anything, nullQuery).Return([]database.Report{
					{
						ID:         reportID1,
						ReportedAt: testTime1,
						ReportedBy: userID,
						Reason:     "Inappropriate content posted last week",
					},
					{
						ID:         reportID2,
						ReportedAt: testTime2,
						ReportedBy: userID,
						Reason:     "New inappropriate content posted today",
					},
				}, nil).Once()

				return reportID1, reportID2, testTime1, testTime2
			},
			validateResp: func(t *testing.T, resp *pb.SearchReportsByDateResponse, reportID1, reportID2 uuid.UUID) {
				assert.NotNil(t, resp)
				assert.Equal(t, 2, len(resp.Report))

				// First report should be the older one
				assert.Equal(t, reportID1.String(), resp.Report[0].Id)
				assert.Contains(t, resp.Report[0].Reason, "last week")

				// Second report should be the newer one
				assert.Equal(t, reportID2.String(), resp.Report[1].Id)
				assert.Contains(t, resp.Report[1].Reason, "today")
			},
		},
		{
			name:  "empty result",
			query: "nonexistent",
			mockSetup: func() (uuid.UUID, uuid.UUID, time.Time, time.Time) {
				reportID1 := uuid.New()
				reportID2 := uuid.New()
				testTime1 := time.Now().Add(-24 * time.Hour)
				testTime2 := time.Now()

				nullQuery := sql.NullString{String: "nonexistent", Valid: true}
				mockDB.On("SearchReportsByDate", mock.Anything, nullQuery).Return([]database.Report{}, nil).Once()

				return reportID1, reportID2, testTime1, testTime2
			},
			validateResp: func(t *testing.T, resp *pb.SearchReportsByDateResponse, reportID1, reportID2 uuid.UUID) {
				assert.NotNil(t, resp)
				assert.Equal(t, 0, len(resp.Report))
			},
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock and get IDs for validation
			reportID1, reportID2, _, _ := tc.mockSetup()

			// Execute the method
			resp, err := testServer.SearchReportsByDate(context.Background(), &pb.SearchReportsByDateRequest{
				Query: tc.query,
			})

			// Validate results
			assert.NoError(t, err)
			tc.validateResp(t, resp, reportID1, reportID2)

			// Verify mock expectations
			mockDB.AssertExpectations(t)
		})
	}
}
