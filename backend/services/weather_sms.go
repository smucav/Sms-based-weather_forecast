package services

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	"github.com/innotechdevops/openmeteo"
	"github.com/paradox-3arthling/africastalking/sms"
	"log"
	//	"net/http"
	//	"net/url"
	"os"
	"strings"
)

// Coordinates holds latitude and longitude
type Coordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// WeatherResponse from Open-Meteo Weather API
type WeatherResponse struct {
	Current struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current_weather"`
}

// Package-level variable for weather codes
var WeatherCodes = map[int]string{
	0:  "Clear sky",
	1:  "Mainly clear",
	2:  "Partly cloudy",
	3:  "Overcast",
	45: "Fog",
	48: "Depositing rime fog",

	51: "Light drizzle",
	53: "Moderate drizzle",
	55: "Dense drizzle",

	56: "Light freezing drizzle",
	57: "Dense freezing drizzle",

	61: "Slight rain",
	63: "Moderate rain",
	65: "Heavy rain",

	66: "Light freezing rain",
	67: "Heavy freezing rain",

	71: "Slight snow fall",
	73: "Moderate snow fall",
	75: "Heavy snow fall",

	77: "Snow grains",

	80: "Slight rain showers",
	81: "Moderate rain showers",
	82: "Violent rain showers",

	85: "Slight snow showers",
	86: "Heavy snow showers",

	95: "Thunderstorm",
	96: "Thunderstorm with slight hail",
	99: "Thunderstorm with heavy hail",
}

type Location struct {
	Region    string
	Woreda    string
	Latitude  float32
	Longitude float32
}

// Package-level variable for location latitude and longitude supports
var knownLocations = []Location{
	{"Addis Ababa", "Arada", 9.036, 38.752},
	{"Addis Ababa", "Kirkos", 9.010, 38.740},
	{"Addis Ababa", "Kolfe Keranio", 9.005, 38.670},
	{"Addis Ababa", "Lideta", 9.030, 38.730},
	{"Addis Ababa", "Nifas Silk-Lafto", 8.980, 38.740},
	{"Addis Ababa", "Yeka", 9.050, 38.780},
	{"Afar", "Asayita", 11.570, 41.440},
	{"Afar", "Awash", 9.010, 40.170},
	{"Afar", "Dubti", 11.730, 41.080},
	{"Afar", "Gewane", 10.150, 40.650},
	{"Afar", "Mille", 11.400, 40.690},
	{"Amhara", "Bahir Dar", 11.600, 37.380},
	{"Amhara", "Debre Markos", 10.340, 37.720},
	{"Amhara", "Debre Tabor", 11.850, 38.016},
	{"Amhara", "Dessie", 11.130, 39.630},
	{"Amhara", "Gondar", 12.600, 37.460},
	{"Amhara", "Kombolcha", 11.080, 39.740},
	{"Benishangul-Gumuz", "Asosa", 10.070, 34.530},
	{"Benishangul-Gumuz", "Bambasi", 9.800, 34.720},
	{"Benishangul-Gumuz", "Dangur", 10.950, 35.350},
	{"Benishangul-Gumuz", "Mao-Komo", 9.300, 34.430},
	{"Benishangul-Gumuz", "Sherkole", 10.150, 34.330},
	{"Dire Dawa", "Gurgura", 9.600, 41.850},
	{"Dire Dawa", "Hara", 9.320, 42.150},
	{"Gambela", "Abobo", 7.850, 34.550},
	{"Gambela", "Gambela", 8.250, 34.590},
	{"Gambela", "Gog", 7.670, 34.350},
	{"Gambela", "Itang", 8.200, 34.290},
	{"Harari", "Amir-Nur", 9.310, 42.120},
	{"Harari", "Abadir", 9.320, 42.130},
	{"Harari", "Shenkor", 9.330, 42.140},
	{"Oromia", "Adama", 8.540, 39.270},
	{"Oromia", "Ambo", 8.980, 37.850},
	{"Oromia", "Bishoftu", 8.750, 38.980},
	{"Oromia", "Jimma", 7.670, 36.830},
	{"Oromia", "Nekemte", 9.090, 36.530},
	{"Oromia", "Shashamane", 7.200, 38.600},
	{"Somali", "Degehabur", 8.220, 43.080},
	{"Somali", "Gode", 5.950, 43.450},
	{"Somali", "Jijiga", 9.350, 42.790},
	{"Somali", "Kebri Beyah", 9.120, 42.250},
	{"Somali", "Warder", 6.970, 45.340},
	{"SNNPR", "Arba Minch", 6.040, 37.550},
	{"SNNPR", "Awasa", 7.060, 38.480},
	{"SNNPR", "Dilla", 6.410, 38.310},
	{"SNNPR", "Sodo", 6.900, 37.760},
	{"SNNPR", "Wolayta", 6.920, 37.650},
	{"Tigray", "Adigrat", 14.280, 39.460},
	{"Tigray", "Axum", 14.120, 38.720},
	{"Tigray", "Mekelle", 13.490, 39.470},
	{"Tigray", "Shire", 14.100, 38.280},
	{"Tigray", "Wukro", 13.790, 39.600},
}

// Used as decode or get the lat and lon of the verbal location
func geocodeLocation(region, woreda string) (Coordinates, error) {
	for _, loc := range knownLocations {
		if strings.EqualFold(loc.Region, region) && strings.EqualFold(loc.Woreda, woreda) {
			return Coordinates{Latitude: loc.Latitude, Longitude: loc.Longitude}, nil
		}
	}
	return Coordinates{}, fmt.Errorf("location not found: %s, %s", region, woreda)
}

func normalizeLocation(raw string) (string, string) {
	// Strip region and zone details like "Oromia", "Arada"
	parts := strings.Split(raw, ",")
	if len(parts) == 0 {
		return raw, raw
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[len(parts)-1])
}

func GetForecast(location string) string {
	zone, woreda := normalizeLocation(location)
	coords, err := geocodeLocation(zone, woreda)
	if err != nil {
		log.Println("Failed to geocode location:", location, err)
		return ""
	}

	// Open-Meteo Weather API endpoint (no API key needed)
	param := openmeteo.Parameter{
		Latitude:       openmeteo.Float32(coords.Latitude),
		Longitude:      openmeteo.Float32(coords.Longitude),
		CurrentWeather: openmeteo.Bool(true),
	}

	m := openmeteo.New()
	resp, err := m.Execute(param)
	//	fmt.Printf("resp: \n", resp)

	if err != nil {
		log.Printf("Failed to fetch weather for %s: %v", location, err)
		return ""
	}

	var weather WeatherResponse

	if err := json.Unmarshal([]byte(resp), &weather); err != nil {
		log.Printf("Failed to parse weather response for %s: %v", location, err)
		return ""
	}

	weatherDesc, ok := WeatherCodes[weather.Current.WeatherCode]
	if !ok {
		weatherDesc = "Unknown weather condition"
	}

	return fmt.Sprintf("Weather update for %s \ntemperature: %.1f°C\nweather: %s\nWind speed: %.1f m/s\n", location, weather.Current.Temperature, weatherDesc, weather.Current.WindSpeed)
}

func SendSMS(phoneNumber, message string) {
	username := os.Getenv("AFRICASTALKING_USERNAME")
	apiKey := os.Getenv("AFRICASTALKING_API_KEY")

	req := sms.Request_data{
		Api_key:  apiKey,
		Username: username,
		To:       []string{phoneNumber},
		Message:  message,
		From:     "GreenHarvest Tech",
	}

	resp, err := req.SendSMS()
	if err != nil {
		log.Printf("Failed to send SMS to %s: %v\n", phoneNumber, err)
		return
	}

	log.Printf("SMS response for %s: %+v\n", phoneNumber, resp)
}

// func SendSMS(phoneNumber, message string) {
// 	username := os.Getenv("AFRICASTALKING_USERNAME")
// 	apiKey := os.Getenv("AFRICASTALKING_API_KEY")

// 	data := url.Values{
// 		"username": {username},
// 		"to":       {phoneNumber},
// 		"message":  {message},
// 		"from":     {"GreenHarvest Tech"},
// 	}

// 	fmt.Println(data)
// 	req, err := http.NewRequest("POST", "https://api.africastalking.com/version1/messaging", bytes.NewBufferString(data.Encode()))
// 	if err != nil {
// 		log.Println("Error creating SMS request:", err)
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("apiKey", apiKey)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println("Error sending SMS:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
// 		log.Printf("SMS sent to %s\n", phoneNumber)
// 	} else {
// 		log.Printf("Failed to send SMS to %s. Status: %s\n", phoneNumber, resp.Status)
// 	}
// }
