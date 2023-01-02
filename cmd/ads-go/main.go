package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	adsgo "github.com/Ivlay/ads-go"
	"github.com/Ivlay/ads-go/internal/pkg/handler"
	"github.com/Ivlay/ads-go/internal/pkg/repository"
	"github.com/Ivlay/ads-go/internal/pkg/service"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to connect DB: %s", err.Error())
	}

	repository := repository.New(db)
	service := service.New(repository)
	handlers := handler.New(service)

	srv := new(adsgo.Server)
	go func() {
		if err := srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			log.Fatal("Error while staring server", err.Error())
		}
	}()

	log.Printf("Server started on the port: %s", viper.GetString("server.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error while server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
