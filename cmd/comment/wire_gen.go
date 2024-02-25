// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/loak155/techbranch-backend/internal/comment/repository"
	"github.com/loak155/techbranch-backend/internal/comment/router"
	"github.com/loak155/techbranch-backend/internal/comment/usecase"
	"github.com/loak155/techbranch-backend/internal/pkg/config"
	"github.com/loak155/techbranch-backend/internal/pkg/db"
	"github.com/loak155/techbranch-backend/proto/comment"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitServer(conf *config.Config, grpcServer *grpc.Server) (comment.CommentServiceServer, error) {
	gormDB := db.NewCommentDB(conf)
	iCommentRepository := repository.NewCommentRepository(gormDB)
	iCommentUsecase := usecase.NewCommentUsecase(iCommentRepository)
	commentServiceServer := router.NewCommentGRPCServer(grpcServer, iCommentUsecase)
	return commentServiceServer, nil
}
