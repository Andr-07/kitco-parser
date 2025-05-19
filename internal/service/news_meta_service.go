package service

import (
	"kitco-parser/internal/models"
	"kitco-parser/internal/parser"
	"kitco-parser/internal/repository"
	"log"
)

type NewsMetaService struct {
	newsMetaRepo *repository.NewsMetaRepository
}

func NewNewsMetaService(repo *repository.NewsMetaRepository) *NewsMetaService {
	return &NewsMetaService{newsMetaRepo: repo}
}

func (s *NewsMetaService) IngestKitcoNews() error {
	newsItems, err := parser.FetchKitcoNews()
	if err != nil {
		return err
	}

	for _, item := range newsItems {
		meta := &models.NewsMeta{
			Title:       item.Title,
			TitleHash:   item.TitleHash,
			URL:         item.URL,
			Source:      item.Source,
			PublishedAt: &item.Published,
			Lang:        item.Lang,
			Status:      "NEW",
		}

		saved, err := s.newsMetaRepo.Save(meta)
		if err != nil {
			return err
		}
		if saved != nil {
			log.Printf("Saved: %s", saved.Title)
		} else {
			log.Printf("Already exists: %s", item.Title)
		}
	}

	return nil
}
