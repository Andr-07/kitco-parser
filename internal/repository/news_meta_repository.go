package repository

import (
	"fmt"
	"kitco-parser/internal/models"
	"kitco-parser/pkg/db"
)

type NewsMetaRepository struct {
	Database *db.Db
}

func NewNewsMetaRepository(database *db.Db) *NewsMetaRepository {
	return &NewsMetaRepository{
		Database: database,
	}
}

func (repo *NewsMetaRepository) Save(item *models.NewsMeta) (*models.NewsMeta, error) {
	var exists bool

	err := repo.Database.
		Model(&models.NewsMeta{}).
		Select("1").
		Where("title_hash = ? AND status = 'NEW'", item.TitleHash).
		Limit(1).
		Scan(&exists).Error
	if err != nil {
		return nil, fmt.Errorf("checking existence: %w", err)
	}

	if exists {
		return nil, nil
	}

	if err := repo.Database.Create(item).Error; err != nil {
		return nil, fmt.Errorf("creating news_meta: %w", err)
	}

	return item, nil
}

