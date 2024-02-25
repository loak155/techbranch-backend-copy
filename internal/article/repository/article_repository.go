package repository

import (
	"github.com/loak155/techbranch-backend/internal/article/domain"
	"gorm.io/gorm"
)

type IArticleRepository interface {
	CreateArticle(article *domain.Article) error
	GetArticle(article *domain.Article, id int) error
	ListArticles(articles *[]domain.Article, offset, limit int) error
	UpdateArticle(article *domain.Article) error
	DeleteArticle(id int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) CreateArticle(article *domain.Article) error {
	if err := ar.db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) GetArticle(article *domain.Article, id int) error {
	if err := ar.db.First(article, id).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) ListArticles(articles *[]domain.Article, offset, limit int) error {
	if err := ar.db.Order("created_at desc").Offset(offset).Limit(limit).Find(articles).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) UpdateArticle(article *domain.Article) error {
	if err := ar.db.Save(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) DeleteArticle(id int) error {
	if err := ar.db.Delete(&domain.Article{}, id).Error; err != nil {
		return err
	}
	return nil
}
