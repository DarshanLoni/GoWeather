package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	openWeatherAPIURL = "http://api.openweathermap.org/data/2.5/weather"
)

// WeatherResponse represents the OpenWeather API response
type WeatherResponse struct {
	Name    string `json:"name"`
	Sys     Sys    `json:"sys"`
	Main    Main   `json:"main"`
	Weather []Weather `json:"weather"`
	Wind    Wind   `json:"wind"`
	Clouds  Clouds `json:"clouds"`
}

type Sys struct {
	Country string `json:"country"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

// WeatherData represents the simplified response to frontend
type WeatherData struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feels_like"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	Pressure    int     `json:"pressure"`
	WindSpeed   float64 `json:"wind_speed"`
	CloudCover  int     `json:"cloud_cover"`
	Icon        string  `json:"icon"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// API routes
	r.HandleFunc("/api/weather", getWeatherHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	fmt.Printf("Weather app running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City parameter is required", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		http.Error(w, "OpenWeather API key not configured", http.StatusInternalServerError)
		return
	}

	weatherData, err := fetchWeatherData(city, apiKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching weather data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}

func fetchWeatherData(city, apiKey string) (*WeatherData, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", openWeatherAPIURL, city, apiKey)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert to simplified format
	weatherData := &WeatherData{
		City:        weatherResp.Name,
		Country:     weatherResp.Sys.Country,
		Temperature: weatherResp.Main.Temp,
		FeelsLike:   weatherResp.Main.FeelsLike,
		Humidity:    weatherResp.Main.Humidity,
		Pressure:    weatherResp.Main.Pressure,
		WindSpeed:   weatherResp.Wind.Speed,
		CloudCover:  weatherResp.Clouds.All,
	}

	if len(weatherResp.Weather) > 0 {
		weatherData.Description = weatherResp.Weather[0].Description
		weatherData.Icon = weatherResp.Weather[0].Icon
	}

	return weatherData, nil
}