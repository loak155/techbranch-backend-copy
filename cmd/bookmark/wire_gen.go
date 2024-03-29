// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/loak155/techbranch-backend/internal/article/client"
	"github.com/loak155/techbranch-backend/internal/bookmark/repository"
	"github.com/loak155/techbranch-backend/internal/bookmark/router"
	"github.com/loak155/techbranch-backend/internal/bookmark/usecase"
	"github.com/loak155/techbranch-backend/internal/pkg/config"
	"github.com/loak155/techbranch-backend/internal/pkg/db"
	"github.com/loak155/techbranch-backend/proto/bookmark"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitServer(conf *config.Config, grpcServer *grpc.Server) (bookmark.BookmarkServiceServer, error) {
	gormDB := db.NewBookmarkDB(conf)
	iBookmarkRepository := repository.NewBookmarkRepository(gormDB)
	articleServiceClient, err := client.NewArticleGRPCClient(conf)
	if err != nil {
		return nil, err
	}
	iBookmarkUsecase := usecase.NewBookmarkUsecase(iBookmarkRepository, articleServiceClient)
	bookmarkServiceServer := router.NewBookmarkGRPCServer(grpcServer, iBookmarkUsecase)
	return bookmarkServiceServer, nil
}
