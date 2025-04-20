package config

import (
	"log"
	"os"
)

func LoadConfig() {
	required := []string{"DATABASE_URL", "AFRICASTALKING_API_KEY", "AFRICASTALKING_USERNAME", "OPENWEATHER_API_KEY"}
	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Missing required environment variable: %s", key)
		}
	}
}
