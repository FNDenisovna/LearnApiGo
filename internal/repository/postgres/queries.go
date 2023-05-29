package postgres

import (
	"LearnApiGo/internal/models"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *Repository) InsertAlbum(album *models.AlbumRow) error {
	query := "INSERT INTO public.album (id, title, artist, price) VALUES (@ID, @Title, @Artist, @Price)"

	/*var args pgx.NamedArgs
	//encoding struct to bytes
	bytesArr, _ := json.Marshal(*album)
	//decoding bytes to map
	json.Unmarshal(bytesArr, &args)*/

	//hand-getting fields
	args := pgx.NamedArgs{
		"ID":     album.Id,
		"Title":  album.Title,
		"Artist": album.Artist,
		"Price":  album.Price,
	}

	_, err := db.db.Exec(context.Background(), query, args)
	if err != nil {
		log.Printf("Error exec inserting into database: %v\n", err)
		return err
	}
	return nil
}

func (db *Repository) SelectAlbums(limit int) (*[]models.AlbumRow, error) {
	var query string

	if limit > 0 {
		query = "SELECT id, title, artist, price FROM public.album ORDER BY title LIMIT @Limit "
	} else {
		query = "SELECT id, title, artist, price FROM public.album ORDER BY Title"
	}

	args := pgx.NamedArgs{
		"Limit": limit,
	}

	rows, err := db.db.Query(context.Background(), query, args)
	if err != nil {
		log.Printf("Error exec getting rows from database: %v\n", err)
		return nil, err
	}
	//res := []models.AlbumRow{}

	res, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.AlbumRow])
	if err != nil {
		log.Printf("Error reading rows from database: %v\n", err)
		return nil, err
	}

	return &res, nil
}

func (db *Repository) SelectAlbum(id *uuid.UUID) (*models.AlbumRow, error) {
	query := "SELECT id, title, artist, price FROM public.album WHERE id = @id"
	args := pgx.NamedArgs{
		"id": id}

	var res models.AlbumRow
	err := db.db.QueryRow(context.Background(), query, args).Scan(&res.Id, &res.Title, &res.Artist, &res.Price)

	if err != nil {
		log.Printf("Error exec getting row from database: %v\n", err)
		return nil, err
	}

	return &res, nil
}
