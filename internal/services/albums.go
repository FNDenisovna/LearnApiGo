package services

import (
	"LearnApiGo/internal/models"

	"github.com/google/uuid"
)

/*var albums = []models.Album{
	{Id: uuid.New(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: uuid.New(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: uuid.New(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}*/

func (s *Service) GetAlbums(limit int) (*[]models.Album, error) {
	rows, err := s.storage.SelectAlbums(limit)
	if err != nil {
		return nil, err
	}

	var res []models.Album

	for _, row := range *rows {
		model, err := row.ToModel()

		if err != nil {
			return nil, err
		}
		res = append(res, *model)
	}
	return &res, nil
}

func (s *Service) GetAlbum(id *uuid.UUID) (*models.Album, error) {
	row, err := s.storage.SelectAlbum(id)
	if err != nil {
		return nil, err
	}

	res, err := row.ToModel()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) CreateAlbum(album *models.Album) error {
	row, err := models.NewRow(album.Title, album.Artist, album.Price)
	if err != nil {
		return err
	}
	if err := s.storage.InsertAlbum(row); err != nil {
		return err
	}

	album.Id = row.Id.String()
	return nil
}
