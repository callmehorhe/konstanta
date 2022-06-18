package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/callmehorhe/konstanta/pkg/handler"
	"github.com/callmehorhe/konstanta/pkg/repository"
	"github.com/callmehorhe/konstanta/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	/* if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	}) */
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal("fail init db")
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	httpServer := &http.Server{
		Addr:           port,
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Print("listen port ", port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
