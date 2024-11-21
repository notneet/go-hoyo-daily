package repository

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"

	"github.com/notneet/go-hoyo-daily/internal/database"
	"github.com/notneet/go-hoyo-daily/internal/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(ctx context.Context, user *model.Article) error
	GetAll(ctx context.Context, limit, offset int) (*ArticleListData, error)
	GetByID(ctx context.Context, id uint64) (*model.Article, error)
	Update(ctx context.Context, updatedArticle *model.Article) error
	Delete(ctx context.Context, id uint64) error
}

type ArticleListData struct {
	Articles    []model.Article
	TotalCount  int
	TotalPages  int
	CurrentPage int
	HasNext     bool
	HasPrev     bool
}

type ArticleRepositoryImpl struct {
	db     *database.DB
	logger *slog.Logger
}

func NewArticleRepository(logger *slog.Logger, db *database.DB) ArticleRepository {
	// Run migrations
	if err := db.AutoMigrate(&model.Article{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		os.Exit(1)
	}

	return &ArticleRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (repo *ArticleRepositoryImpl) Create(ctx context.Context, article *model.Article) error {
	err := repo.db.DB.WithContext(ctx).Create(article).Error
	repo.logger.Info("created article", "article", article)

	return err
}

func (repo *ArticleRepositoryImpl) GetAll(ctx context.Context, limit, offset int) (*ArticleListData, error) {
	var articles []model.Article
	if err := repo.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		repo.logger.Error("failed to get all articles", "error", err)

		return nil, err
	}

	var totalCount int64
	if err := repo.db.WithContext(ctx).Model(&model.Article{}).Count(&totalCount).Error; err != nil {
		repo.logger.Error("failed to count articles", "error", err)

		return nil, err
	}

	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))
	currentPage := offset/limit + 1
	hasNext := currentPage < totalPages
	hasPrev := currentPage > 1

	return &ArticleListData{
		Articles:    articles,
		TotalCount:  int(totalCount),
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		HasNext:     hasNext,
		HasPrev:     hasPrev,
	}, nil
}

func (repo *ArticleRepositoryImpl) GetByID(ctx context.Context, id uint64) (*model.Article, error) {
	var article model.Article
	if err := repo.db.WithContext(ctx).First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		repo.logger.Error("failed to get article by id", "error", err)

		return nil, err
	}

	return &article, nil
}

func (repo *ArticleRepositoryImpl) Update(ctx context.Context, updatedArticle *model.Article) error {
	if err := repo.db.WithContext(ctx).Save(updatedArticle).Error; err != nil {
		repo.logger.Error("failed to update article", "error", err)

		return err
	}

	return nil
}

func (repo *ArticleRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	if err := repo.db.WithContext(ctx).Delete(&model.Article{}, id).Error; err != nil {
		repo.logger.Error("failed to delete article", "error", err)

		return err
	}

	return nil
}
