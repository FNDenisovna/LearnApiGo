package main

//swagger https://github.com/swaggo/swag#general-api-info
import (
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"LearnApiGo/docs"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	//https://github.com/MartinHeinz/go-project-blueprint/blob/rest-api/cmd/blueprint/main.go
	//пример подключения бд, авторизация

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"} //, "https"

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	//Используйте эту Run функцию, чтобы подключить маршрутизатор к http.Serverсерверу и запустить его.
	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbums godoc
// @Summary Retrieves all albums
// @Produce json
// @Success 200 {object} jsonresult.JSONResult{data=[]album} "desc"
// @Router /albums [get]
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums) //c.IndentedJSON(http.StatusOK, albums) //c.JSON(albums)
}
