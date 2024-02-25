package repository

import (
	"github.com/loak155/techbranch-backend/internal/bookmark/domain"
	"gorm.io/gorm"
)

type IBookmarkRepository interface {
	CreateBookmark(bookmark *domain.Bookmark) error
	UpdateBookmarkWithUnscoped(id int) error
	GetBookmark(bookmark *domain.Bookmark, id int) error
	GetBookmarkByUserIDAndArticleID(bookmark *domain.Bookmark, userID, articleID int) error
	GetBookmarkByUserIDAndArticleIDWithUnscoped(bookmark *domain.Bookmark, userID, articleID int) error
	ListBookmarks(bookmarks *[]domain.Bookmark) error
	ListBookmarksByUserID(bookmarks *[]domain.Bookmark, userID int) error
	ListBookmarksByArticleID(bookmarks *[]domain.Bookmark, articleID int) error
	DeleteBookmark(id int) error
	DeleteBookmarkByUserIDAndArticleID(userID, articleID int) error
	DeleteBookmarkByUserID(UserID int) error
	UpdateBookmarkByUserIDWithUnscoped(UserID int) error
	DeleteBookmarkByArticleID(ArticleID int) error
	UpdateBookmarkByArticleIDWithUnscoped(ArticleID int) error
}

type bookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) IBookmarkRepository {
	return &bookmarkRepository{db}
}

func (br *bookmarkRepository) CreateBookmark(bookmark *domain.Bookmark) error {
	if err := br.db.Create(bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) UpdateBookmarkWithUnscoped(id int) error {
	if err := br.db.Unscoped().Model(&domain.Bookmark{}).Where(id).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) GetBookmark(bookmark *domain.Bookmark, id int) error {
	if err := br.db.First(bookmark, id).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) GetBookmarkByUserIDAndArticleID(bookmark *domain.Bookmark, userID, articleID int) error {
	if err := br.db.Where("user_id=? AND article_id=?", userID, articleID).First(bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) GetBookmarkByUserIDAndArticleIDWithUnscoped(bookmark *domain.Bookmark, userID, articleID int) error {
	if err := br.db.Unscoped().Where("user_id=? AND article_id=?", userID, articleID).First(bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) ListBookmarks(bookmarks *[]domain.Bookmark) error {
	if err := br.db.Find(bookmarks).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) ListBookmarksByUserID(bookmarks *[]domain.Bookmark, userID int) error {
	if err := br.db.Where("user_id=?", userID).Find(bookmarks).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) ListBookmarksByArticleID(bookmarks *[]domain.Bookmark, articleID int) error {
	if err := br.db.Where("user_id=?", articleID).Find(bookmarks).Error; err != nil {
		return err
	}
	return nil

}

func (br *bookmarkRepository) DeleteBookmark(id int) error {
	if err := br.db.Delete(&domain.Bookmark{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) DeleteBookmarkByUserIDAndArticleID(userID, articleID int) error {
	if err := br.db.Where("user_id=? AND article_id=?", userID, articleID).Delete(&domain.Bookmark{}).Error; err != nil {
		return err
	}
	return nil
}

func (cr *bookmarkRepository) DeleteBookmarkByUserID(userID int) error {
	if err := cr.db.Delete(&domain.Bookmark{}, "user_id=?", userID).Error; err != nil {
		return err
	}
	return nil
}

func (cr *bookmarkRepository) UpdateBookmarkByUserIDWithUnscoped(userID int) error {
	if err := cr.db.Unscoped().Model(&domain.Bookmark{}).Where("user_id", userID).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (cr *bookmarkRepository) DeleteBookmarkByArticleID(articleID int) error {
	if err := cr.db.Delete(&domain.Bookmark{}, "article_id=?", articleID).Error; err != nil {
		return err
	}
	return nil
}

func (cr *bookmarkRepository) UpdateBookmarkByArticleIDWithUnscoped(articleID int) error {
	if err := cr.db.Unscoped().Model(&domain.Bookmark{}).Where("article_id", articleID).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	return nil
}
