package postgres

import (
	"LearnApiGo/internal/models"

	"github.com/google/uuid"
)

type IStorage interface {
	InsertAlbum(album *models.AlbumRow) error
	SelectAlbums(limit int) (*[]models.AlbumRow, error)
	SelectAlbum(id *uuid.UUID) (*models.AlbumRow, error)
	SelectUser(login string) ([]byte, error)
}
