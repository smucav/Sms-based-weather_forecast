package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"sms-weather-backend/config"
	"sms-weather-backend/database"
	"sms-weather-backend/handlers"
	"sms-weather-backend/scheduler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	config.LoadConfig()
	database.Init()
	scheduler.StartScheduler()

	http.HandleFunc("/register", handlers.RegisterUser)

	http.HandleFunc("/trigger-weather", scheduler.TriggerWeatherCheck)

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
