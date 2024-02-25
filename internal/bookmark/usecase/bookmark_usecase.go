package usecase

import (
	"context"

	"github.com/loak155/techbranch-backend/internal/bookmark/domain"
	"github.com/loak155/techbranch-backend/internal/bookmark/repository"
	pb "github.com/loak155/techbranch-backend/proto/article"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IBookmarkUsecase interface {
	CreateBookmark(ctx context.Context, bookmark domain.Bookmark) (domain.Bookmark, error)
	GetBookmark(ctx context.Context, id int) (domain.Bookmark, error)
	ListBookmarks(ctx context.Context) ([]domain.Bookmark, error)
	ListBookmarksByUserID(ctx context.Context, UserID int) ([]domain.Bookmark, error)
	ListBookmarksByArticleID(ctx context.Context, articleID int) ([]domain.Bookmark, error)
	DeleteBookmark(ctx context.Context, bookmark domain.Bookmark) (bool, error)
	DeleteBookmarkByUserID(ctx context.Context, UserID int) (bool, error)
	DeleteBookmarkByUserIDCompensate(ctx context.Context, UserID int) (bool, error)
	DeleteBookmarkByArticleID(ctx context.Context, ArticleID int) (bool, error)
	DeleteBookmarkByArticleIDCompensate(ctx context.Context, ArticleID int) (bool, error)
}

type bookmarkUsecase struct {
	br repository.IBookmarkRepository
	ac pb.ArticleServiceClient
}

func NewBookmarkUsecase(br repository.IBookmarkRepository, ac pb.ArticleServiceClient) IBookmarkUsecase {
	return &bookmarkUsecase{br, ac}
}

func (bu *bookmarkUsecase) CreateBookmark(ctx context.Context, bookmark domain.Bookmark) (domain.Bookmark, error) {
	_, err := bu.ac.IncrementBookmarksCount(ctx, &pb.IncrementBookmarksCountRequest{Id: int32(bookmark.ArticleID)})
	if err != nil {
		return domain.Bookmark{}, err
	}

	storeBookmark := domain.Bookmark{}
	err = bu.br.GetBookmarkByUserIDAndArticleIDWithUnscoped(&storeBookmark, int(bookmark.UserID), int(bookmark.ArticleID))
	if storeBookmark.ID != 0 && err == nil {
		if err := bu.br.UpdateBookmarkWithUnscoped(int(storeBookmark.ID)); err != nil {
			bu.ac.IncrementBookmarksCountCompensate(ctx, &pb.IncrementBookmarksCountRequest{Id: int32(bookmark.ArticleID)})
			return domain.Bookmark{}, status.Errorf(codes.Internal, "failed to compensate bookmark deletion: %v", err)
		}
		return storeBookmark, nil
	} else {
		if err := bu.br.CreateBookmark(&bookmark); err != nil {
			bu.ac.IncrementBookmarksCountCompensate(ctx, &pb.IncrementBookmarksCountRequest{Id: int32(bookmark.ArticleID)})
			return domain.Bookmark{}, status.Errorf(codes.Internal, "failed to create bookmark: %v", err)
		}
		return bookmark, nil
	}

}

func (bu *bookmarkUsecase) GetBookmark(ctx context.Context, id int) (domain.Bookmark, error) {
	storedBookmark := domain.Bookmark{}
	if err := bu.br.GetBookmark(&storedBookmark, id); err != nil {
		return domain.Bookmark{}, status.Errorf(codes.Internal, "failed to get bookmark: %v", err)
	}
	return storedBookmark, nil
}

func (bu *bookmarkUsecase) ListBookmarks(ctx context.Context) ([]domain.Bookmark, error) {
	storedBookmarks := []domain.Bookmark{}
	if err := bu.br.ListBookmarks(&storedBookmarks); err != nil {
		return []domain.Bookmark{}, status.Errorf(codes.Internal, "failed to get bookmark list: %v", err)
	}
	return storedBookmarks, nil

}

func (bu *bookmarkUsecase) ListBookmarksByUserID(ctx context.Context, UserID int) ([]domain.Bookmark, error) {
	storedBookmark := []domain.Bookmark{}
	if err := bu.br.ListBookmarksByUserID(&storedBookmark, UserID); err != nil {
		return []domain.Bookmark{}, status.Errorf(codes.Internal, "failed to get bookmark list: %v", err)
	}
	return storedBookmark, nil

}

func (bu *bookmarkUsecase) ListBookmarksByArticleID(ctx context.Context, articleID int) ([]domain.Bookmark, error) {
	storedBookmark := []domain.Bookmark{}
	if err := bu.br.ListBookmarksByArticleID(&storedBookmark, articleID); err != nil {
		return []domain.Bookmark{}, status.Errorf(codes.Internal, "failed to get bookmark list: %v", err)
	}
	return storedBookmark, nil
}

func (bu *bookmarkUsecase) DeleteBookmark(ctx context.Context, bookmark domain.Bookmark) (bool, error) {
	_, err := bu.ac.DecrementBookmarksCount(ctx, &pb.DecrementBookmarksCountRequest{Id: int32(bookmark.ArticleID)})
	if err != nil {
		return false, err
	}

	if err := bu.br.DeleteBookmarkByUserIDAndArticleID(int(bookmark.UserID), int(bookmark.ArticleID)); err != nil {
		bu.ac.DecrementBookmarksCountCompensate(ctx, &pb.DecrementBookmarksCountRequest{Id: int32(bookmark.ArticleID)})
		return false, status.Errorf(codes.Internal, "failed to delete bookmark: %v", err)
	}

	return true, nil
}

func (cu *bookmarkUsecase) DeleteBookmarkByUserID(ctx context.Context, userID int) (bool, error) {
	if err := cu.br.DeleteBookmarkByUserID(userID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to delete bookmark: %v", err)
	}
	return true, nil
}

func (cu *bookmarkUsecase) DeleteBookmarkByUserIDCompensate(ctx context.Context, userID int) (bool, error) {
	if err := cu.br.UpdateBookmarkByUserIDWithUnscoped(userID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to compensate bookmark deletion: %v", err)
	}
	return true, nil
}

func (cu *bookmarkUsecase) DeleteBookmarkByArticleID(ctx context.Context, articleID int) (bool, error) {
	if err := cu.br.DeleteBookmarkByArticleID(articleID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to delete bookmark: %v", err)
	}
	return true, nil
}

func (cu *bookmarkUsecase) DeleteBookmarkByArticleIDCompensate(ctx context.Context, articleID int) (bool, error) {
	if err := cu.br.UpdateBookmarkByArticleIDWithUnscoped(articleID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to compensate bookmark deletion: %v", err)
	}
	return true, nil
}
