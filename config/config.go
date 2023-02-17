package config

import (
	"LearnApiGo/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DatabaseUrl = "postgres://postgres:postgres@localhost:5432/apigo"

var DbPool pgxpool.Pool

// TODO Singltone pgxpool
//https://donchev.is/post/working-with-postgresql-in-go-using-pgx/

func ConnectDb() {

	db, err := pgxpool.New(context.Background(), DatabaseUrl)
	DbPool = *db

	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
	}

	//defer DbPool.Close()
	log.Println("Successfully connected to database")
}

func Close() {
	DbPool.Close()
}

func InsertAlbum(album *models.Album) *error {
	query := "INSERT INTO public.album (ID, Title, Artist, Price) VALUES (@ID, @Title, @Artist, @Price)"

	args := pgx.NamedArgs{
		"ID":     album.ID,
		"Title":  album.Title,
		"Artist": album.Artist,
		"Price":  album.Price,
	}
	log.Printf("In InsertAlbum getting args")

	_, err := DbPool.Exec(context.Background(), query, args)
	log.Printf("In InsertAlbum getting exec")
	if err != nil {
		log.Printf("Error exec inserting into database: %v\n", err)
		return &err
	}
	return nil
}
