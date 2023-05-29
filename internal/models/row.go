package models

import (
	"github.com/google/uuid"
)

type AlbumRow struct {
	Id     uuid.UUID
	Title  string
	Artist string
	Price  float64
}

func (row *AlbumRow) ToModel() (*Album, error) {
	return &Album{
		Id:     row.Id.String(),
		Title:  row.Title,
		Artist: row.Artist,
		Price:  row.Price,
	}, nil
}
