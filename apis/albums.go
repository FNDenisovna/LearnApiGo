package apis

import (
	"LearnApiGo/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// PostAlbums godoc
// @Summary Adds an album to db from JSON received in the request body
// @Produce json
// @Param request body []models.Album true "query params"//
// @Success 200 {object} string
// @Router /albums [post]
func PostAlbums(c *gin.Context) {
	var err = services.PostAlbums(c)
	if err == nil {
		c.IndentedJSON(http.StatusCreated, "Ok")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
}

// GetAlbums godoc
// @Summary Retrieves all albums
// @Produce json
// @Success 200 {object} []models.Album
// @Router /albums [get]
func GetAlbums(c *gin.Context) {
	var albums = services.GetAlbums()
	c.JSON(http.StatusOK, albums) //c.IndentedJSON(http.StatusOK, albums) //c.JSON(albums)
}

// GetAlbum godoc
// @Summary Retrieves album with ID
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} models.Album
// @Router /album/{id} [get]
func GetAlbum(c *gin.Context) {
	id := c.Param("id")
	album := services.GetAlbum(&id)
	if album == nil {
		//c.AbortWithStatus(http.StatusNotFound, "Where is no album with getting ID in db.")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.JSON(http.StatusOK, album)
	}
}
