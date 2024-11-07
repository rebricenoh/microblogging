package main

import (
	"log"
	"net/http"

	"microblogging/internal/domain"
	"microblogging/internal/handler"
	"microblogging/internal/repository"
	"microblogging/internal/service"
	"microblogging/pkg/config"
)

func main() {
	// Configurar el entorno
	cfg := config.NewConfig()

	// Crear repositorios
	tweetRepo := repository.NewGormTweetRepository(cfg.DB)
	//followRepo := repository.NewGormFollowRepository(cfg.DB)

	// Crear servicios
	tweetService := service.NewTweetService(tweetRepo)
	//followService := service.NewFollowService(followRepo)

	// Crear handlers
	tweetHandler := handler.NewTweetHandler(tweetService)
	// Puedes agregar un handler para followService en el futuro si es necesario

	// Rutas de la API
	http.HandleFunc("/tweet", tweetHandler.PostTweetHandler)

	// Migración de la base de datos
	cfg.DB.AutoMigrate(&domain.Tweet{}, &domain.Follow{})

	log.Println("Servidor iniciado en :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
