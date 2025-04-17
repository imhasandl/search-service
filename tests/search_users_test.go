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

func TestSearchUsers(t *testing.T) {
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
		validateResp   func(t *testing.T, resp *pb.SearchUsersResponse)
	}{
		{
			name:  "successful search",
			query: "john",
			mockSetup: func() {
				userID := uuid.New()
				nullQuery := sql.NullString{String: "john", Valid: true}

				mockDB.On("SearchUsers", mock.Anything, nullQuery).Return([]database.User{
					{
						ID:                     userID,
						CreatedAt:              testTime,
						UpdatedAt:              testTime,
						Email:                  "john@example.com",
						Password:               "hashedpassword",
						Username:               "johndoe",
						Subscribers:            []uuid.UUID{},
						SubscribedTo:           []uuid.UUID{},
						IsPremium:              true,
						VerificationCode:       12345,
						VerificationExpireTime: testTime.Add(24 * time.Hour),
						IsVerified:             true,
					},
				}, nil).Once()
			},
			expectedError: false,
			validateResp: func(t *testing.T, resp *pb.SearchUsersResponse) {
				assert.NotNil(t, resp)
				assert.Equal(t, 1, len(resp.Users))
				assert.Equal(t, "johndoe", resp.Users[0].Username)
				assert.Equal(t, true, resp.Users[0].IsPremium)
				assert.Equal(t, true, resp.Users[0].IsVerified)
			},
		},
		{
			name:  "empty query",
			query: "",
			mockSetup: func() {
				nullQuery := sql.NullString{String: "", Valid: false}
				mockDB.On("SearchUsers", mock.Anything, nullQuery).Return([]database.User{}, nil).Once()
			},
			expectedError: false,
			validateResp: func(t *testing.T, resp *pb.SearchUsersResponse) {
				assert.NotNil(t, resp)
				assert.Equal(t, 0, len(resp.Users))
			},
		},
		{
			name:  "database error",
			query: "error",
			mockSetup: func() {
				nullQuery := sql.NullString{String: "error", Valid: true}
				mockDB.On("SearchUsers", mock.Anything, nullQuery).Return(
					[]database.User{}, errors.New("database error"),
				).Once()
			},
			expectedError:  true,
			expectedCode:   codes.Internal,
			expectedErrMsg: "can't get users",
			validateResp:   nil,
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock for this test case
			tc.mockSetup()

			// Execute the method
			resp, err := testServer.SearchUsers(context.Background(), &pb.SearchUsersRequest{
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

func TestSearchUsersByDate(t *testing.T) {
	// Test setup
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")

	// Define test cases
	testCases := []struct {
		name         string
		query        string
		mockSetup    func() (uuid.UUID, uuid.UUID, time.Time, time.Time)
		validateResp func(t *testing.T, resp *pb.SearchUsersByDateResponse, userID1, userID2 uuid.UUID)
	}{
		{
			name:  "successful search with sorting",
			query: "user",
			mockSetup: func() (uuid.UUID, uuid.UUID, time.Time, time.Time) {
				userID1 := uuid.New()
				userID2 := uuid.New()
				testTime1 := time.Now().Add(-24 * time.Hour) // older
				testTime2 := time.Now()                      // newer

				nullQuery := sql.NullString{String: "user", Valid: true}
				mockDB.On("SearchUsersByDate", mock.Anything, nullQuery).Return([]database.User{
					{
						ID:         userID1,
						CreatedAt:  testTime1,
						UpdatedAt:  testTime1,
						Email:      "user1@example.com",
						Username:   "user1",
						IsPremium:  false,
						IsVerified: true,
					},
					{
						ID:         userID2,
						CreatedAt:  testTime2,
						UpdatedAt:  testTime2,
						Email:      "user2@example.com",
						Username:   "user2",
						IsPremium:  true,
						IsVerified: false,
					},
				}, nil).Once()

				return userID1, userID2, testTime1, testTime2
			},
			validateResp: func(t *testing.T, resp *pb.SearchUsersByDateResponse, userID1, userID2 uuid.UUID) {
				assert.NotNil(t, resp)
				assert.Equal(t, 2, len(resp.Users))

				// First user should be the older one
				assert.Equal(t, userID1.String(), resp.Users[0].Id)
				assert.Equal(t, "user1", resp.Users[0].Username)

				// Second user should be the newer one
				assert.Equal(t, userID2.String(), resp.Users[1].Id)
				assert.Equal(t, "user2", resp.Users[1].Username)
			},
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock and get IDs for validation
			userID1, userID2, _, _ := tc.mockSetup()

			// Execute the method
			resp, err := testServer.SearchUsersByDate(context.Background(), &pb.SearchUsersByDateRequest{
				Query: tc.query,
			})

			// Validate results
			assert.NoError(t, err)
			tc.validateResp(t, resp, userID1, userID2)

			// Verify mock expectations
			mockDB.AssertExpectations(t)
		})
	}
}
