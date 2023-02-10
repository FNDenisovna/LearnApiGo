package services

import (
	"LearnApiGo/models"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Get all info albums from db
func GetAlbums() (arr *[]models.Album) {
	return &albums
}

// Get album with gettin ID
func GetAlbum(id *string) (album *models.Album) {
	for _, a := range albums {
		if a.ID == *id {
			return &a
		}
	}
	return nil
}

// Add new row into db
func PostAlbums(c *gin.Context) *error {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return &err
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	return nil
}
