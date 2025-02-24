package server

import (
	"context"
	"database/sql"

	"github.com/imhasandl/search-service/internal/database"
	pb "github.com/imhasandl/search-service/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	searchUserParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	users, err := s.db.SearchUsers(ctx, searchUserParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get users: %v - SearchUser", err)
	}

	responseUsers := make([]*pb.User, len(users))
	for i, user := range users {
		responseUsers[i] = &pb.User{
			Id:               user.ID.String(),
			CreatedAt:        timestamppb.New(user.CreatedAt),
			UpdatedAt:        timestamppb.New(user.UpdatedAt),
			Email:            user.Email,
			Username:         user.Username,
			IsPremium:        user.IsPremium,
			VerificationCode: user.VerificationCode,
			IsVerified:       user.IsVerified,
		}
	}
	return &pb.SearchUsersResponse{
		Users: responseUsers,
	}, nil
}

func (s *server) SearchPosts(ctx context.Context, req *pb.SearchPostsRequest) (*pb.SearchPostsResponse, error) {
	searchPostParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	posts, err := s.db.SearchPosts(ctx, searchPostParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get posts: %v - SearchUser", err)
	}

	responsePosts := make([]*pb.Post, len(posts))
	for i, post := range posts {
		responsePosts[i] = &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy,
			Body:      post.Body,
			Views:     post.Views,
			LikedBy:   post.LikedBy,
		}
	}

	return &pb.SearchPostsResponse{
		Post: responsePosts,
	}, nil
}

func (s *server) SearchReports(ctx context.Context, req *pb.SearchReportsRequest) (*pb.SearchReportsResponse, error) {
	searchReportsParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	reports, err := s.db.SearchReports(ctx, searchReportsParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get report: %v - SearchReports", err)
	}

	responseReports := make([]*pb.Report, len(reports))
	for i, report := range reports {
		responseReports[i] = &pb.Report{
			Id:        report.ID.String(),
			ReportedAt: timestamppb.New(report.ReportedAt),
			ReportedBy: report.ReportedBy.String(),
			Reason: report.Reason,
		}
	}

	return &pb.SearchReportsResponse{
		Report: responseReports,
	}, nil
}

func (s *server) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	return &pb.CheckResponse{
		Status: req.GetCheck(),
	}, nil
}