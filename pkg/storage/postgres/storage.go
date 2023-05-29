package postgres

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Store struct {
	Pool *pgxpool.Pool
}

var (
	pgInstance *Store
	pgOnce     sync.Once
)

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

func initDefaultEnv() error {
	if len(os.Getenv("PGHOST")) == 0 {
		if err := os.Setenv("PGHOST", "localhost"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGPORT")) == 0 {
		if err := os.Setenv("PGPORT", "5432"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGDATABASE")) == 0 {
		if err := os.Setenv("PGDATABASE", "apigo"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGUSER")) == 0 {
		if err := os.Setenv("PGUSER", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGPASSWORD")) == 0 {
		if err := os.Setenv("PGPASSWORD", "postgres"); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(os.Getenv("PGSSLMODE")) == 0 {
		if err := os.Setenv("PGSSLMODE", "disable"); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func toDSN(s Settings) string {
	var args []string

	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	}

	if s.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	}

	if len(s.Database) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", s.Database))
	}

	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", s.User))
	}

	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	}

	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	}

	return strings.Join(args, " ")
}

func New(settings Settings) (*Store, error) {
	config, err := pgxpool.ParseConfig(toDSN(settings))
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &Store{Pool: conn}, nil
}

func (pg *Store) Ping(ctx context.Context) error {
	return pg.Pool.Ping(ctx)
}

func (pg *Store) Close() {
	pg.Pool.Close()
}

// TODO Singltone pgxpool
//https://donchev.is/post/working-with-postgresql-in-go-using-pgx/

/*
var DatabaseUrl = "postgres://postgres:postgres@localhost:5432/apigo"

var DbPool pgxpool.Pool

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
*/
