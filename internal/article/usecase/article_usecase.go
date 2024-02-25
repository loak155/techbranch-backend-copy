package usecase

import (
	"context"

	"github.com/loak155/techbranch-backend/internal/article/domain"
	"github.com/loak155/techbranch-backend/internal/article/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IArticleUsecase interface {
	CreateArticle(ctx context.Context, article domain.Article) (domain.Article, error)
	GetArticle(ctx context.Context, id int) (domain.Article, error)
	ListArticles(ctx context.Context, offset, limit int) ([]domain.Article, error)
	UpdateArticle(ctx context.Context, article domain.Article) (bool, error)
	DeleteArticle(ctx context.Context, id int) (bool, error)
	IncrementBookmarksCount(ctx context.Context, id int) (bool, error)
	IncrementBookmarksCountCompensate(ctx context.Context, id int) (bool, error)
	DecrementBookmarksCount(ctx context.Context, id int) (bool, error)
	DecrementBookmarksCountCompensate(ctx context.Context, id int) (bool, error)
}

type articleUsecase struct {
	ar repository.IArticleRepository
}

func NewArticleUsecase(ar repository.IArticleRepository) IArticleUsecase {
	return &articleUsecase{ar}
}

func (uu *articleUsecase) CreateArticle(ctx context.Context, article domain.Article) (domain.Article, error) {
	article.BookmarkCount = 0
	if err := uu.ar.CreateArticle(&article); err != nil {
		return domain.Article{}, status.Errorf(codes.Internal, "failed to create article: %v", err)
	}
	return article, nil
}

func (uu *articleUsecase) GetArticle(ctx context.Context, id int) (domain.Article, error) {
	storedArticle := domain.Article{}
	if err := uu.ar.GetArticle(&storedArticle, id); err != nil {
		return domain.Article{}, status.Errorf(codes.Internal, "failed to get article: %v", err)
	}
	return storedArticle, nil
}

func (uu *articleUsecase) ListArticles(ctx context.Context, offset, limit int) ([]domain.Article, error) {
	articles := []domain.Article{}
	if err := uu.ar.ListArticles(&articles, offset, limit); err != nil {
		return []domain.Article{}, status.Errorf(codes.Internal, "failed to get article list: %v", err)
	}
	return articles, nil
}

func (uu *articleUsecase) UpdateArticle(ctx context.Context, article domain.Article) (bool, error) {
	updatedArticle := domain.Article{
		ID:            article.ID,
		Title:         article.Title,
		Url:           article.Url,
		BookmarkCount: article.BookmarkCount,
	}
	if err := uu.ar.UpdateArticle(&updatedArticle); err != nil {
		return false, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}
	return true, nil
}

func (uu *articleUsecase) DeleteArticle(ctx context.Context, id int) (bool, error) {
	if err := uu.ar.DeleteArticle(id); err != nil {
		return false, status.Errorf(codes.Internal, "failed to delete article: %v", err)
	}
	return true, nil
}

func (uu *articleUsecase) IncrementBookmarksCount(ctx context.Context, id int) (bool, error) {
	article := domain.Article{}
	if err := uu.ar.GetArticle(&article, id); err != nil {
		return false, status.Errorf(codes.NotFound, "failed to get article: %v", err)
	}
	article.BookmarkCount++
	if err := uu.ar.UpdateArticle(&article); err != nil {
		return false, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}
	return true, nil
}

func (uu *articleUsecase) IncrementBookmarksCountCompensate(ctx context.Context, id int) (bool, error) {
	article := domain.Article{}
	if err := uu.ar.GetArticle(&article, id); err != nil {
		return false, status.Errorf(codes.NotFound, "failed to get article: %v", err)
	}
	article.BookmarkCount--
	if err := uu.ar.UpdateArticle(&article); err != nil {
		return false, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}
	return true, nil
}

func (uu *articleUsecase) DecrementBookmarksCount(ctx context.Context, id int) (bool, error) {
	article := domain.Article{}
	if err := uu.ar.GetArticle(&article, id); err != nil {
		return false, status.Errorf(codes.NotFound, "failed to get article: %v", err)
	}
	article.BookmarkCount--
	if err := uu.ar.UpdateArticle(&article); err != nil {
		return false, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}
	return true, nil
}

func (uu *articleUsecase) DecrementBookmarksCountCompensate(ctx context.Context, id int) (bool, error) {
	article := domain.Article{}
	if err := uu.ar.GetArticle(&article, id); err != nil {
		return false, status.Errorf(codes.NotFound, "failed to get article: %v", err)
	}
	article.BookmarkCount++
	if err := uu.ar.UpdateArticle(&article); err != nil {
		return false, status.Errorf(codes.Internal, "failed to update article: %v", err)
	}
	return true, nil
}
