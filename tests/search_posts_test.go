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

func TestSearchPosts(t *testing.T) {
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
		validateResp   func(t *testing.T, resp *pb.SearchPostsResponse)
	}{
		{
			name:  "successful search",
			query: "hello",
			mockSetup: func() {
				postID := uuid.New()
				userID := uuid.New()
				nullQuery := sql.NullString{String: "hello", Valid: true}

				mockDB.On("SearchPosts", mock.Anything, nullQuery).Return([]database.Post{
					{
						ID:        postID,
						CreatedAt: testTime,
						UpdatedAt: testTime,
						PostedBy:  userID,
						Body:      "Hello world post",
						Views:     42,
						Likes:     10,
						LikedBy:   []string{"user1", "user2"},
					},
				}, nil).Once()
			},
			expectedError: false,
			validateResp: func(t *testing.T, resp *pb.SearchPostsResponse) {
				assert.NotNil(t, resp)
				assert.Equal(t, 1, len(resp.Post))
				assert.Contains(t, resp.Post[0].Body, "Hello")
				assert.Equal(t, int32(42), resp.Post[0].Views)
				assert.Equal(t, int32(10), resp.Post[0].Likes)
				assert.Equal(t, 2, len(resp.Post[0].LikedBy))
			},
		},
		{
			name:  "empty query",
			query: "",
			mockSetup: func() {
				nullQuery := sql.NullString{String: "", Valid: false}
				mockDB.On("SearchPosts", mock.Anything, nullQuery).Return([]database.Post{}, nil).Once()
			},
			expectedError: false,
			validateResp: func(t *testing.T, resp *pb.SearchPostsResponse) {
				assert.NotNil(t, resp)
				assert.Equal(t, 0, len(resp.Post))
			},
		},
		{
			name:  "database error",
			query: "error",
			mockSetup: func() {
				nullQuery := sql.NullString{String: "error", Valid: true}
				mockDB.On("SearchPosts", mock.Anything, nullQuery).Return(
					[]database.Post{}, errors.New("database error"),
				).Once()
			},
			expectedError:  true,
			expectedCode:   codes.Internal,
			expectedErrMsg: "can't find posts by date",
			validateResp:   nil,
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock for this test case
			tc.mockSetup()

			// Execute the method
			resp, err := testServer.SearchPosts(context.Background(), &pb.SearchPostsRequest{
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

func TestSearchPostsByDate(t *testing.T) {
	// Test setup
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")

	// Define test cases
	testCases := []struct {
		name         string
		query        string
		mockSetup    func() (uuid.UUID, uuid.UUID, time.Time, time.Time)
		validateResp func(t *testing.T, resp *pb.SearchPostsByDateResponse, postID1, postID2 uuid.UUID)
	}{
		{
			name:  "successful search with sorting",
			query: "post",
			mockSetup: func() (uuid.UUID, uuid.UUID, time.Time, time.Time) {
				postID1 := uuid.New()
				postID2 := uuid.New()
				userID := uuid.New()
				testTime1 := time.Now().Add(-24 * time.Hour) // older
				testTime2 := time.Now()                      // newer

				nullQuery := sql.NullString{String: "post", Valid: true}
				mockDB.On("SearchPostsByDate", mock.Anything, nullQuery).Return([]database.Post{
					{
						ID:        postID1,
						CreatedAt: testTime1,
						UpdatedAt: testTime1,
						PostedBy:  userID,
						Body:      "First post content",
						Views:     10,
						Likes:     5,
						LikedBy:   []string{"user1"},
					},
					{
						ID:        postID2,
						CreatedAt: testTime2,
						UpdatedAt: testTime2,
						PostedBy:  userID,
						Body:      "Second post content",
						Views:     20,
						Likes:     15,
						LikedBy:   []string{"user1", "user2", "user3"},
					},
				}, nil).Once()

				return postID1, postID2, testTime1, testTime2
			},
			validateResp: func(t *testing.T, resp *pb.SearchPostsByDateResponse, postID1, postID2 uuid.UUID) {
				assert.NotNil(t, resp)
				assert.Equal(t, 2, len(resp.Post))

				// First post should be the older one
				assert.Equal(t, postID1.String(), resp.Post[0].Id)
				assert.Equal(t, "First post content", resp.Post[0].Body)
				assert.Equal(t, int32(10), resp.Post[0].Views)
				assert.Equal(t, 1, len(resp.Post[0].LikedBy))

				// Second post should be the newer one
				assert.Equal(t, postID2.String(), resp.Post[1].Id)
				assert.Equal(t, "Second post content", resp.Post[1].Body)
				assert.Equal(t, int32(15), resp.Post[1].Likes)
				assert.Equal(t, 3, len(resp.Post[1].LikedBy))
			},
		},
		{
			name:  "empty result",
			query: "nonexistent",
			mockSetup: func() (uuid.UUID, uuid.UUID, time.Time, time.Time) {
				postID1 := uuid.New()
				postID2 := uuid.New()
				testTime1 := time.Now().Add(-24 * time.Hour)
				testTime2 := time.Now()

				nullQuery := sql.NullString{String: "nonexistent", Valid: true}
				mockDB.On("SearchPostsByDate", mock.Anything, nullQuery).Return([]database.Post{}, nil).Once()

				return postID1, postID2, testTime1, testTime2
			},
			validateResp: func(t *testing.T, resp *pb.SearchPostsByDateResponse, postID1, postID2 uuid.UUID) {
				assert.NotNil(t, resp)
				assert.Equal(t, 0, len(resp.Post))
			},
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock and get IDs for validation
			postID1, postID2, _, _ := tc.mockSetup()

			// Execute the method
			resp, err := testServer.SearchPostsByDate(context.Background(), &pb.SearchPostsByDateRequest{
				Query: tc.query,
			})

			// Validate results
			assert.NoError(t, err)
			tc.validateResp(t, resp, postID1, postID2)

			// Verify mock expectations
			mockDB.AssertExpectations(t)
		})
	}
}
