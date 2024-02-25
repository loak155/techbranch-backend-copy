package repository

import (
	"github.com/loak155/techbranch-backend/internal/comment/domain"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *domain.Comment) error
	GetComment(comment *domain.Comment, id int) error
	ListCommentsByArticleID(comment *[]domain.Comment, articleID int) error
	UpdateComment(comment *domain.Comment) error
	DeleteComment(id int) error
	DeleteCommentByUserID(userID int) error
	UpdateByUserIDWithUnscoped(userID int) error
	DeleteCommentByArticleID(articleID int) error
	UpdateByArticleIDWithUnscoped(articleID int) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(comment *domain.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) GetComment(comment *domain.Comment, id int) error {
	if err := cr.db.First(comment, id).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) ListCommentsByArticleID(comments *[]domain.Comment, articleID int) error {
	if err := cr.db.Where("article_id=?", articleID).Find(comments).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) UpdateComment(comment *domain.Comment) error {
	if err := cr.db.Save(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) DeleteComment(id int) error {
	if err := cr.db.Delete(&domain.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) DeleteCommentByUserID(userID int) error {
	if err := cr.db.Delete(&domain.Comment{}, "user_id=?", userID).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) UpdateByUserIDWithUnscoped(userID int) error {
	if err := cr.db.Unscoped().Model(&domain.Comment{}).Where("user_id", userID).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) DeleteCommentByArticleID(articleID int) error {
	if err := cr.db.Delete(&domain.Comment{}, "article_id=?", articleID).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) UpdateByArticleIDWithUnscoped(articleID int) error {
	if err := cr.db.Unscoped().Model(&domain.Comment{}).Where("article_id", articleID).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}
