package models

import (
	"net/http"

	"github.com/google/uuid"
)

type Album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (a *Album) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *Album) Bind(r *http.Request) error {
	return nil
}

func NewRow(title string, artist string, price float64) (*AlbumRow, error) {
	//Business conditions

	return &AlbumRow{
		Id:     uuid.New(),
		Title:  title,
		Artist: artist,
		Price:  price,
	}, nil
}
