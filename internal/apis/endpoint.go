package apis

import (
	"LearnApiGo/internal/models"

	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// PostAlbums godoc
// @Summary Adds an album to db from JSON received in the request body
// @Produce json
// @Param request body models.Album true "query params"//
// @Success 200 {object} models.Album
// @Router /album/create [post]
func (api *Api) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	newAlbum := &models.Album{}

	if err := render.Bind(r, newAlbum); err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err, models.ParseErr+"albumModel", 400))
		return
	}

	var err = api.service.CreateAlbum(newAlbum)
	if err == nil {
		render.Status(r, http.StatusCreated)
		render.Render(w, r, newAlbum)
	} else {
		render.Render(w, r, models.ErrInvalidRequest(err, "Not created", 500))
	}
}

// GetAlbums godoc
// @Summary Retrieves all albums
// @Produce json
// @Param limit path int true "Limit"
// @Success 200 {object} []models.Album
// @Router /albums/{limit} [get]
func (api *Api) GetAlbums(w http.ResponseWriter, r *http.Request) {
	limit := r.Context().Value("limit").(int)

	albums, err := api.service.GetAlbums(limit)
	if err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err, models.GetListErr, 500))
	} else {
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, NewAlbumListResponse(albums))
	}
}

func NewAlbumListResponse(albums *[]models.Album) []render.Renderer {
	list := []render.Renderer{}
	for _, album := range *albums {
		list = append(list, &album)
	}
	return list
}

// GetAlbum godoc
// @Summary Retrieves album with ID
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} models.Album
// @Router /album/{id} [get]
func (api *Api) GetAlbum(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(uuid.UUID)

	album, err := api.service.GetAlbum(&id)
	if err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err, models.GetItemErr, 500))
	} else if album == nil {
		//c.AbortWithStatus(http.StatusNotFound, "Where is no album with getting ID in db.")
		render.Render(w, r, models.ErrInvalidRequest(err, models.NotFoundErr, 404))
	} else {
		render.Status(r, http.StatusOK)
		render.Render(w, r, album)
	}
}
