package services

import (
	"LearnApiGo/internal/models"
	storage "LearnApiGo/internal/repository/postgres"

	"github.com/google/uuid"
)

type IAlbums interface {
	GetAlbums(limit int) (*[]models.Album, error)
	GetAlbum(id *uuid.UUID) (*models.Album, error)
	CreateAlbum(album *models.Album) error
}

type Service struct {
	storage storage.IStorage
}

func New(storage storage.IStorage) *Service {
	var s = &Service{
		storage: storage,
	}
	return s
}
