package apis

import (
	"LearnApiGo/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums godoc
// @Summary Adds an album to db from JSON received in the request body
// @Produce json
// @Success 200 {object} jsonresult.JSONResult{data=string}
// @Router /albums [post]
func PostAlbums(c *gin.Context) {
	var err = services.PostAlbums(c)
	if err == nil {
		c.IndentedJSON(http.StatusCreated, "Ok")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
}

// getAlbums godoc
// @Summary Retrieves all albums
// @Produce json
// @Success 200 {object} jsonresult.JSONResult{data=[]album} "desc"
// @Router /albums [get]
func GetAlbums(c *gin.Context) {
	var albums = services.GetAlbums()
	c.JSON(http.StatusOK, albums) //c.IndentedJSON(http.StatusOK, albums) //c.JSON(albums)
}
