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

func TestSearchPosts(t *testing.T) {
	// Setup
	mockDB := mocks.NewMockQueries()
	testServer := server.NewServer(mockDB, "test-secret")

	t.Run("successful search", func(t *testing.T) {
		// Prepare test data
		postID := uuid.New()
		userID := uuid.New()
		testTime := time.Now()
		query := "hello"
		nullQuery := sql.NullString{String: query, Valid: true}

		// Configure mock
		mockDB.On("SearchPosts", mock.Anything, nullQuery).Return([]database.Post{
			{
				ID:        postID,
				CreatedAt: testTime,
				UpdatedAt: testTime,
				PostedBy:  userID,
				Body:      "Hello world post content",
				Likes:     5,
				Views:     10,
				LikedBy:   []string{"user1", "user2"},
			},
		}, nil).Once()

		// Execute the method
		resp, err := testServer.SearchPosts(context.Background(), &pb.SearchPostsRequest{
			Query: query,
		})

		// Verify
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 1, len(resp.Post))
		assert.Equal(t, postID.String(), resp.Post[0].Id)
		assert.Equal(t, userID.String(), resp.Post[0].PostedBy)
		assert.Equal(t, "Hello world post content", resp.Post[0].Body)
		assert.Equal(t, int32(5), resp.Post[0].Likes)
		assert.Equal(t, int32(10), resp.Post[0].Views)
		assert.Equal(t, []string{"user1", "user2"}, resp.Post[0].LikedBy)

		// Verify mock expectations
		mockDB.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		// Configure mock to return error
		query := "error"
		nullQuery := sql.NullString{String: query, Valid: true}

		mockDB.On("SearchPosts", mock.Anything, nullQuery).Return(
			[]database.Post{}, errors.New("database error"),
		).Once()

		// Execute the method
		resp, err := testServer.SearchPosts(context.Background(), &pb.SearchPostsRequest{
			Query: query,
		})

		// Verify error is returned correctly
		assert.Error(t, err)
		assert.Nil(t, resp)

		// Check if error is correctly formatted as gRPC error
		statusErr, ok := status.FromError(err)
		assert.True(t, ok)
		assert.Equal(t, codes.Internal, statusErr.Code())
		assert.Contains(t, statusErr.Message(), "can't find posts")

		// Verify mock expectations
		mockDB.AssertExpectations(t)
	})
}

func TestSearchPostsByDate(t *testing.T) {

}
