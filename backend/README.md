# 🌦️ SMS Weather Alert Backend

A backend service written in Go that sends daily weather forecast updates to users via SMS, using the [Africa's Talking API](https://africastalking.com/) and [Open-Meteo Weather API](https://open-meteo.com/). Designed for farmers and informal workers who rely on timely weather alerts but lack access to smartphones or the internet.

## 📦 Features

- 🛰️ Fetches real-time weather data using Open-Meteo.
- 📱 Sends SMS weather alerts using Africa's Talking.
- 🗓️ Daily scheduler to dispatch forecasts automatically.
- 🌍 Supports geolocation-based lookup (region → lat/lon).
- 🧾 PostgreSQL-backed user registration and storage.
- 🔘 Manual weather check trigger via API endpoint.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **SMS Provider:** Africa's Talking
- **Weather Provider:** Open-Meteo
- **Scheduler:** `time.Ticker` (Go concurrency)
- **API Wrapper:** [`github.com/paradox-3arthling/africastalking`](https://github.com/paradox-3arthling/africastalking)

## 📂 Project Structure

```
sms-weather-backend/
├── database/            # DB connection and init logic
├── models/              # User model definition
├── scheduler/           # Scheduled weather dispatch logic
├── services/            # Weather & SMS integration logic
├── weather/             # Open-Meteo API integration
├── utils/               # Geolocation helpers, weather code mapping
├── main.go              # Entry point and router setup
├── go.mod
└── README.md
```

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/sms-weather-backend.git
cd sms-weather-backend
```

### 2. Set Up Environment Variables

Create a `.env` file or set these in your environment:

```bash
AFRICASTALKING_USERNAME=your_username
AFRICASTALKING_API_KEY=your_api_key
DATABASE_URL=postgres://user:pass@localhost:5432/weatherdb
```

### 3. Run Database Migrations

> Ensure your PostgreSQL instance is up and running.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    location TEXT NOT NULL
);
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Server

```bash
go run main.go
```

## 🧪 Manual Weather Trigger API

```
POST /trigger-weather
```

Use this to manually dispatch forecasts outside the daily scheduler. Useful for testing.

## 💬 Sample SMS Output

```
Hello, your weather forecast for today:
Location: Mekelle
Temperature: 26.4°C
Windspeed: 12.3 km/h
Condition: Partly Cloudy
```

## 🧠 Weather Code Reference

```go
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
  61: "Light rain",
  63: "Moderate rain",
  65: "Heavy rain",
  71: "Light snow",
  73: "Moderate snow",
  75: "Heavy snow",
  80: "Rain showers",
  81: "Moderate showers",
  82: "Violent rain showers",
  95: "Thunderstorm",
}
```

## 🤝 Contributions

PRs, issues, and suggestions are welcome. Help grow this to support multilingual alerts, regional customizations, and more intelligent scheduling.

## 📧 Contact

Built with ❤️ by GreenHarvest Tech Team
Email: [smucav@gmail.com]