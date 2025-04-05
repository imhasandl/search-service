package server

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/imhasandl/search-service/cmd/helper"
	"github.com/imhasandl/search-service/internal/database"
	pb "github.com/imhasandl/search-service/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DatabaseQuerier defines the interface for database operations used by the search service.
// It contains methods for searching users, posts, and reports with various filtering options.
type DatabaseQuerier interface {
	SearchUsers(ctx context.Context, arg sql.NullString) ([]database.User, error)
	SearchUsersByDate(ctx context.Context, arg sql.NullString) ([]database.User, error)
	SearchPosts(ctx context.Context, arg sql.NullString) ([]database.Post, error)
	SearchPostsByDate(ctx context.Context, arg sql.NullString) ([]database.Post, error)
	SearchReports(ctx context.Context, arg sql.NullString) ([]database.Report, error)
	SearchReportsByDate(ctx context.Context, arg sql.NullString) ([]database.Report, error)
}

// Server represents the gRPC server for the search service.
type Server interface {
	pb.SearchServiceServer
}

type server struct {
	pb.UnimplementedSearchServiceServer
	db          DatabaseQuerier
	tokenSecret string
}

// NewServer creates and returns a new instance of the search service server.
// It requires database queries implementation and a token secret for authentication.
func NewServer(dbQueries DatabaseQuerier, tokenSecret string) Server {
	return &server{
		db:          dbQueries,
		tokenSecret: tokenSecret,
	}
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	startTime := time.Now()
	defer func() {
		endTime := time.Since(startTime)
		log.Printf("Finished searching in %v", endTime)
	}()

	searchUserParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	users, err := s.db.SearchUsers(ctx, searchUserParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get users - SearchUsers", err)
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

func (s *server) SearchUsersByDate(ctx context.Context, req *pb.SearchUsersByDateRequest) (*pb.SearchUsersByDateResponse, error) {
	searchUsersByDateParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	users, err := s.db.SearchUsersByDate(ctx, searchUsersByDateParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get users by date", err)
	}

	responseUsersByDate := make([]*pb.User, len(users))
	for i, user := range users {
		responseUsersByDate[i] = &pb.User{
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

	return &pb.SearchUsersByDateResponse{
		Users: responseUsersByDate,
	}, nil
}

func (s *server) SearchPosts(ctx context.Context, req *pb.SearchPostsRequest) (*pb.SearchPostsResponse, error) {
	searchPostParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	posts, err := s.db.SearchPosts(ctx, searchPostParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't find posts by date - SearchPostsByDate", err)
	}

	responsePosts := make([]*pb.Post, len(posts))
	for i, post := range posts {
		responsePosts[i] = &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy.String(),
			Body:      post.Body,
			Likes:     post.Likes,
			Views:     post.Views,
			LikedBy:   post.LikedBy,
		}
	}

	return &pb.SearchPostsResponse{
		Post: responsePosts,
	}, nil
}

func (s *server) SearchPostsByDate(ctx context.Context, req *pb.SearchPostsByDateRequest) (*pb.SearchPostsByDateResponse, error) {
	searchPostsByDateParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	posts, err := s.db.SearchPostsByDate(ctx, searchPostsByDateParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get posts by date - SearchPostsByDate", err)
	}

	responsePostsByDate := make([]*pb.Post, len(posts))
	for i, post := range posts {
		responsePostsByDate[i] = &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy.String(),
			Body:      post.Body,
			Likes:     post.Likes,
			Views:     post.Views,
			LikedBy:   post.LikedBy,
		}
	}

	return &pb.SearchPostsByDateResponse{
		Post: responsePostsByDate,
	}, nil
}

func (s *server) SearchReports(ctx context.Context, req *pb.SearchReportsRequest) (*pb.SearchReportsResponse, error) {
	searchReportsParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	reports, err := s.db.SearchReports(ctx, searchReportsParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get report - SearchReports", err)
	}

	responseReports := make([]*pb.Report, len(reports))
	for i, report := range reports {
		responseReports[i] = &pb.Report{
			Id:         report.ID.String(),
			ReportedAt: timestamppb.New(report.ReportedAt),
			ReportedBy: report.ReportedBy.String(),
			Reason:     report.Reason,
		}
	}

	return &pb.SearchReportsResponse{
		Report: responseReports,
	}, nil
}

func (s *server) SearchReportsByDate(ctx context.Context, req *pb.SearchReportsByDateRequest) (*pb.SearchReportsByDateResponse, error) {
	searchReportsByDateParams := sql.NullString{String: req.GetQuery(), Valid: req.GetQuery() != ""}

	reports, err := s.db.SearchReportsByDate(ctx, searchReportsByDateParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get posts by date - SearchPostsByDate", err)
	}

	responseReports := make([]*pb.Report, len(reports))
	for i, report := range reports {
		responseReports[i] = &pb.Report{
			Id:         report.ID.String(),
			ReportedAt: timestamppb.New(report.ReportedAt),
			ReportedBy: report.ReportedBy.String(),
			Reason:     report.Reason,
		}
	}

	return &pb.SearchReportsByDateResponse{
		Report: responseReports,
	}, nil
}
