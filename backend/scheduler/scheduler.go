package scheduler

import (
	"log"
	"net/http"
	"time"

	"sms-weather-backend/database"
	"sms-weather-backend/models"
	"sms-weather-backend/services"
)

func StartScheduler() {
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		for {
			checkAndSendWeather()
			<-ticker.C
		}
	}()
}

func checkAndSendWeather() {
	rows, err := database.DB.Query("SELECT name, phone_number, location FROM users")
	if err != nil {
		log.Println("Error querying users:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Name, &u.PhoneNumber, &u.Location)
		if err != nil {
			log.Println("Error scanning user:", err)
			continue
		}

		forecast := services.GetForecast(u.Location)
		if forecast != "" {
			services.SendSMS(u.PhoneNumber, forecast)
		}
	}
}

// function to trigger weather check manually
func TriggerWeatherCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	checkAndSendWeather() // Reuse the existing function

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Weather check triggered successfully"))
}
