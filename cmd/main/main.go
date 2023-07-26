package main

//swagger https://github.com/swaggo/swag#general-api-info
import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"LearnApiGo/internal/apis"
	"LearnApiGo/internal/grpc"
	repo "LearnApiGo/internal/repository/postgres"
	services "LearnApiGo/internal/services"
	storage "LearnApiGo/pkg/storage/postgres"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("SERVICE_NAME", "albumService")
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("GRPC_PORT", 8000)
}

func main() {

	//Подключение к БД
	conn, err := storage.New(storage.Settings{})
	if err != nil {
		panic(err)
	}

	db := repo.New(conn.Pool)
	service := services.New(db)
	go apis.New(service)
	go grpc.New(service)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	defer conn.Pool.Close()
}
