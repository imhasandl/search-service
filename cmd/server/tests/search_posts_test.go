package tests

// import (
// 	"context"
// 	"log"
// 	"testing"

// 	"github.com/imhasandl/search-service/cmd/server"
// 	"github.com/imhasandl/search-service/cmd/server/tests/mocks"
// 	pb "github.com/imhasandl/search-service/protos"
// 	"google.golang.org/grpc/codes"
// )

// func TestSearchPosts(t *testing.T) {
// 	// Test setup
// 	mockDB := mocks.NewMockQueries()
// 	testServer := server.NewServer(mockDB, "test-secret")

// 	// Define test cases
// 	testcase := []struct {
// 		name           string
// 		query          string
// 		mockSetup      func()
// 		expectedError  bool
// 		expectedCode   codes.Code
// 		expectedErrMsg string
// 		validateResp   func(t *testing.T, resp *pb.SearchUsersResponse)
// 	}{
// 		{},
// 	}


// 	// Execute test cases
// 	for _, tc := range testcase {
// 		t.Run(tc.name, func(t *testing.T) {

// 			resp, _ := testServer.SearchPosts(context.Background(), &pb.SearchPostsRequest{
// 				Query: tc.query,
// 			})

// 			log.Print(resp)
// 		})
// 	}
// }

// func TestSearchPostsByDate(t *testing.T) {

// }
