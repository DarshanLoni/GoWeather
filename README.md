# GoWeather 🌤️

A simple and elegant weather application built with Go that fetches weather data from the OpenWeather API and displays it through a beautiful web interface.

## Features

- 🌍 Get current weather for any city worldwide
- 🎨 Clean and responsive web interface
- 📱 Mobile-friendly design
- ⚡ Fast API responses with Go backend
- 🌡️ Detailed weather information including:
  - Current temperature and "feels like" temperature
  - Weather description and icon
  - Humidity and atmospheric pressure
  - Wind speed and cloud coverage
- 📄 Environment configuration with .env file support

## Prerequisites

- Go 1.21 or higher
- OpenWeather API key (free at [openweathermap.org](https://openweathermap.org/api))

## Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/DarshanLoni/GoWeather.git
cd GoWeather
```

### 2. Get an OpenWeather API key
- Visit [OpenWeatherMap](https://openweathermap.org/api)
- Sign up for a free account
- Generate an API key

### 3. Configure environment variables
Create a `.env` file in the root directory (or copy from `.env.example`):
```env
OPENWEATHER_API_KEY=your_api_key_here
PORT=8081
```

### 4. Install dependencies
```bash
go mod tidy
```

### 5. Run the application

**For Windows:**
```cmd
"C:\Program Files\Go\bin\go.exe" run main.go
```

**For Linux/macOS:**
```bash
go run main.go
```

### 6. Open your browser
Visit: `http://localhost:8081`

## Project Structure

```
GoWeather/
├── main.go              # Main application server
├── go.mod              # Go module file
├── go.sum              # Go dependencies checksum
├── .env                # Environment variables (create this)
├── .env.example        # Environment variables template
├── static/             # Static web assets
│   ├── index.html      # Main HTML page
│   ├── style.css       # Stylesheet
│   └── script.js       # Frontend JavaScript
└── README.md           # This file
```

## API Endpoints

- `GET /` - Serves the main web interface
- `GET /api/weather?city={city_name}` - Returns weather data for the specified city

### Example API Response

```json
{
  "city": "London",
  "country": "GB",
  "temperature": 15.5,
  "feels_like": 14.2,
  "description": "partly cloudy",
  "humidity": 72,
  "pressure": 1013,
  "wind_speed": 3.2,
  "cloud_cover": 40,
  "icon": "02d"
}
```

## Technologies Used

- **Backend:** Go with Gorilla Mux router
- **Frontend:** Vanilla HTML, CSS, and JavaScript
- **API:** OpenWeather API
- **Environment:** godotenv for .env file support
- **Styling:** CSS Grid, Flexbox, and CSS animations

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `OPENWEATHER_API_KEY` | Your OpenWeather API key | Required |
| `PORT` | Server port number | 8080 |

## Features Explained

### Backend (Go)
- RESTful API design with proper error handling
- Clean JSON responses
- Environment variable configuration with .env file support
- HTTP client with timeout
- Modular code structure

### Frontend
- Responsive design that works on all devices
- Loading states and error handling
- Smooth animations and transitions
- Clean, modern UI with glassmorphism effects
- Accessibility considerations

## Development

To run in development mode:
```bash
# Windows
"C:\Program Files\Go\bin\go.exe" run main.go

# Linux/macOS
go run main.go
```

To build the application:
```bash
# Windows
"C:\Program Files\Go\bin\go.exe" build -o weather-app.exe

# Linux/macOS
go build -o weather-app
```

## Customization

You can easily customize the application:

- **Styling:** Modify `static/style.css` to change colors, fonts, or layout
- **Features:** Add more weather data points by updating the structs in `main.go`
- **UI:** Enhance the frontend by modifying `static/index.html` and `static/script.js`
- **Configuration:** Update `.env` file for different ports or settings

## Troubleshooting

**API Key Issues:**
- Make sure your API key is correctly set in the `.env` file
- Verify your API key is active on OpenWeatherMap

**City Not Found:**
- Check the spelling of the city name
- Try using the format "City, Country" for better results

**Port Issues:**
- If port 8081 is in use, change the `PORT` value in your `.env` file
- Make sure no other applications are using the same port

**Network Issues:**
- Ensure you have an active internet connection
- Check if the OpenWeather API is accessible from your network

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.

## Author

- **Darshan Loni** - [GitHub](https://github.com/DarshanLoni)

## Acknowledgments

- OpenWeatherMap for providing the weather API
- Gorilla Mux for the excellent HTTP router
- The Go community for amazing tools and libraries