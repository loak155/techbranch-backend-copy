package usecase

import (
	"context"

	"github.com/loak155/techbranch-backend/internal/comment/domain"
	"github.com/loak155/techbranch-backend/internal/comment/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ICommentUsecase interface {
	CreateComment(ctx context.Context, comment domain.Comment) (domain.Comment, error)
	GetComment(ctx context.Context, id int) (domain.Comment, error)
	ListCommentsByArticleID(ctx context.Context, articleID int) ([]domain.Comment, error)
	UpdateComment(ctx context.Context, comment domain.Comment) (bool, error)
	DeleteComment(ctx context.Context, id int) (bool, error)
	DeleteCommentByUserID(ctx context.Context, userID int) (bool, error)
	DeleteCommentByUserIDCompensate(ctx context.Context, userID int) (bool, error)
	DeleteCommentByArticleID(ctx context.Context, articleID int) (bool, error)
	DeleteCommentByArticleIDCompensate(ctx context.Context, articleID int) (bool, error)
}

type commentUsecase struct {
	cr repository.ICommentRepository
}

func NewCommentUsecase(ur repository.ICommentRepository) ICommentUsecase {
	return &commentUsecase{ur}
}

func (cu *commentUsecase) CreateComment(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	if err := cu.cr.CreateComment(&comment); err != nil {
		return domain.Comment{}, status.Errorf(codes.Internal, "failed to create comment: %v", err)
	}
	return comment, nil
}

func (cu *commentUsecase) GetComment(ctx context.Context, id int) (domain.Comment, error) {
	storedComment := domain.Comment{}
	if err := cu.cr.GetComment(&storedComment, id); err != nil {
		return domain.Comment{}, status.Errorf(codes.Internal, "failed to get comment: %v", err)
	}
	return storedComment, nil
}

func (cu *commentUsecase) ListCommentsByArticleID(ctx context.Context, articleID int) ([]domain.Comment, error) {
	storedComment := []domain.Comment{}
	if err := cu.cr.ListCommentsByArticleID(&storedComment, articleID); err != nil {
		return []domain.Comment{}, status.Errorf(codes.Internal, "failed to get comment list: %v", err)
	}
	return storedComment, nil
}

func (cu *commentUsecase) UpdateComment(ctx context.Context, comment domain.Comment) (bool, error) {
	storedComment := domain.Comment{}
	if err := cu.cr.GetComment(&storedComment, int(comment.ID)); err != nil {
		return false, status.Errorf(codes.Internal, "failed to get comment: %v", err)
	}

	storedComment.Content = comment.Content

	if err := cu.cr.UpdateComment(&storedComment); err != nil {
		return false, status.Errorf(codes.Internal, "failed to update comment: %v", err)
	}
	return true, nil
}

func (cu *commentUsecase) DeleteComment(ctx context.Context, id int) (bool, error) {
	if err := cu.cr.DeleteComment(id); err != nil {
		return false, status.Errorf(codes.Internal, "failed to delete comment: %v", err)
	}
	return true, nil
}

func (cu *commentUsecase) DeleteCommentByUserID(ctx context.Context, userID int) (bool, error) {
	if err := cu.cr.DeleteCommentByUserID(userID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to delete comment: %v", err)
	}
	return true, nil
}

func (cu *commentUsecase) DeleteCommentByUserIDCompensate(ctx context.Context, userID int) (bool, error) {
	if err := cu.cr.UpdateByUserIDWithUnscoped(userID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to compensate comment deletion: %v", err)
	}
	return true, nil
}

func (cu *commentUsecase) DeleteCommentByArticleID(ctx context.Context, articleID int) (bool, error) {
	if err := cu.cr.DeleteCommentByArticleID(articleID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to delete comment: %v", err)
	}
	return true, nil
}

func (cu *commentUsecase) DeleteCommentByArticleIDCompensate(ctx context.Context, articleID int) (bool, error) {
	if err := cu.cr.UpdateByArticleIDWithUnscoped(articleID); err != nil {
		return false, status.Errorf(codes.Internal, "failed to compensate comment deletion: %v", err)
	}
	return true, nil
}
