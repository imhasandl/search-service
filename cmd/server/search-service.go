package server

import (
	"context"

	"github.com/imhasandl/search-service/internal/database"
	pb "github.com/imhasandl/search-service/protos"
)

type server struct {
	pb.UnimplementedSearchServiceServer
	db          *database.Queries
	tokenSecret string
}

func NewServer(dbQueries *database.Queries, tokenSecret string) *server {
	return &server{
		db:          dbQueries,
		tokenSecret: tokenSecret,
	}
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	return nil, nil
}

func (s *server) SearchPosts(ctx context.Context, req *pb.SearchPostsRequest) (*pb.SearchPostsResponse, error) {
	return nil, nil
}

