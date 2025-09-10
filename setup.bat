@echo off
echo Weather App Setup Script
echo ========================

echo.
echo Checking if Go is installed...
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo ERROR: Go is not installed or not in PATH
    echo.
    echo Please install Go from: https://golang.org/dl/
    echo After installation, restart your terminal and run this script again.
    echo.
    pause
    exit /b 1
)

echo Go found! Version:
go version

echo.
echo Installing dependencies...
go mod tidy

if %errorlevel% neq 0 (
    echo ERROR: Failed to install dependencies
    pause
    exit /b 1
)

echo.
echo Dependencies installed successfully!
echo.
echo IMPORTANT: Before running the app, set your OpenWeather API key:
echo.
echo Run this command in PowerShell:
echo $env:OPENWEATHER_API_KEY="your_api_key_here"
echo.
echo Or in Command Prompt:
echo set OPENWEATHER_API_KEY=your_api_key_here
echo.
echo Get your free API key from: https://openweathermap.org/api
echo.
echo After setting the API key, run: go run main.go
echo Then open: http://localhost:8080
echo.
pause