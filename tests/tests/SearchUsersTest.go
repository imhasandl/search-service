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

func TestSearchUsers(t *testing.T) {
	// Setup
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")

	t.Run("successful search", func(t *testing.T) {
		// Prepare Test Data
		userID := uuid.New()
		testTime := time.Now()
		query := "john"
		nullQuery := sql.NullString{String: query, Valid: true}

		// Configure Mock
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

		// Execute the method
		resp, err := testServer.SearchUsers(context.Background(), &pb.SearchUsersRequest{
			Query: query,
		})

		// Verify
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 1, len(resp.Users))
		assert.Equal(t, userID.String(), resp.Users[0].Id)
		assert.Equal(t, "johndoe", resp.Users[0].Username)
		assert.Equal(t, true, resp.Users[0].IsPremium)
		assert.Equal(t, true, resp.Users[0].IsVerified)

		// Verify mock expectations
		mockDB.AssertExpectations(t)
	})

	t.Run("empty query", func(t *testing.T) {
		// Empty query should still work but might return all users
		nullQuery := sql.NullString{String: "", Valid: true}

		// Configure mock
		mockDB.On("SearchUsers", mock.Anything, nullQuery).Return([]database.User{}, nil).Once()

		// Execute the method
		resp, err := testServer.SearchUsers(context.Background(), &pb.SearchUsersRequest{
			Query: "",
		})

		// Verify
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 0, len(resp.Users))

		// Verify mock expectations
		mockDB.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		// Configure mock to return error
		query := "error"
		nullQuery := sql.NullString{String: query, Valid: true}

		mockDB.On("SearchUsers", mock.Anything, nullQuery).Return(
			[]database.User{}, errors.New("database error"),
		).Once()

		// Execute the method
		resp, err := testServer.SearchUsers(context.Background(), &pb.SearchUsersRequest{
			Query: query,
		})

		// Verify error is returned correctly
		assert.Error(t, err)
		assert.Nil(t, resp)

		// Check if error is correctly formatted as gRPC error
		statusErr, ok := status.FromError(err)
		assert.True(t, ok)
		assert.Equal(t, codes.Internal, statusErr.Code())
		assert.Contains(t, statusErr.Message(), "can't get users")

		// Verify mock expectations
		mockDB.AssertExpectations(t)
	})
}

func TestSearchUsersByDate(t *testing.T) {
	// Setup
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")
	
	t.Run("successful search with sorting", func(t *testing.T) {
		 // Prepare test data
		 userID1 := uuid.New()
		 userID2 := uuid.New()
		 testTime1 := time.Now().Add(-24 * time.Hour) // older
		 testTime2 := time.Now()                     // newer
		 query := "user"
		 nullQuery := sql.NullString{String: query, Valid: true}
		 
		 // Configure mock to return sorted results
		 mockDB.On("SearchUsersByDate", mock.Anything, nullQuery).Return([]database.User{
			  {
					ID:        userID1,
					CreatedAt: testTime1,
					UpdatedAt: testTime1,
					Email:     "user1@example.com",
					Username:  "user1",
					IsPremium: false,
					IsVerified: true,
			  },
			  {
					ID:        userID2,
					CreatedAt: testTime2,
					UpdatedAt: testTime2,
					Email:     "user2@example.com",
					Username:  "user2",
					IsPremium: true,
					IsVerified: false,
			  },
		 }, nil).Once()
		 
		 // Execute the method
		 resp, err := testServer.SearchUsersByDate(context.Background(), &pb.SearchUsersByDateRequest{
			  Query: query,
		 })
		 
		 // Verify
		 assert.NoError(t, err)
		 assert.NotNil(t, resp)
		 assert.Equal(t, 2, len(resp.Users))
		 
		 // First user should be the older one
		 assert.Equal(t, userID1.String(), resp.Users[0].Id)
		 assert.Equal(t, "user1", resp.Users[0].Username)
		 
		 // Second user should be the newer one
		 assert.Equal(t, userID2.String(), resp.Users[1].Id)
		 assert.Equal(t, "user2", resp.Users[1].Username)
		 
		 // Verify mock expectations
		 mockDB.AssertExpectations(t)
	})
}