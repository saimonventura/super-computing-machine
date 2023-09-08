package main

import (
	"net/http"

	"super-computing-machine/auth/application"
	"super-computing-machine/auth/infrastructure/config"
	"super-computing-machine/auth/infrastructure/dbseed"
	"super-computing-machine/auth/infrastructure/initializers"
	"super-computing-machine/auth/infrastructure/logger"
	"super-computing-machine/auth/infrastructure/persistence/redis"
	"super-computing-machine/auth/interfaces"

	"github.com/gorilla/mux"
)

var authHandler *interfaces.AuthHandler

func main() {
	log := logger.NewLogrusLogger()
	natsMessenger := initializers.InitNats()

	// Initialize Redis connection
	rdb := redis.InitializeRedis()

	// Initialize development user
	seeder := dbseed.NewSeeder(rdb, log)
	seeder.InitializeDevUser()

	authService := application.NewAuthService(rdb, config.GetEnvWithDefault("JWT_SECRET_KEY", "sua_chave_secreta"), natsMessenger)
	authHandler = interfaces.NewAuthHandler(authService)

	r := mux.NewRouter()
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	http.Handle("/", r)

	port := config.GetEnvWithDefault("AUTH_PORT", "8081")

	log.Info("Server started on port " + port)
	log.Debug(http.ListenAndServe(":"+port, r))
}
