document.addEventListener('DOMContentLoaded', function() {
    const cityInput = document.getElementById('cityInput');
    const searchBtn = document.getElementById('searchBtn');
    const loadingSpinner = document.getElementById('loadingSpinner');
    const errorMessage = document.getElementById('errorMessage');
    const weatherCard = document.getElementById('weatherCard');

    // Elements for weather data
    const cityName = document.getElementById('cityName');
    const countryName = document.getElementById('countryName');
    const temperature = document.getElementById('temperature');
    const feelsLike = document.getElementById('feelsLike');
    const weatherDescription = document.getElementById('weatherDescription');
    const humidity = document.getElementById('humidity');
    const pressure = document.getElementById('pressure');
    const windSpeed = document.getElementById('windSpeed');
    const cloudCover = document.getElementById('cloudCover');
    const weatherIcon = document.getElementById('weatherIcon');
    const errorText = document.getElementById('errorText');

    // Event listeners
    searchBtn.addEventListener('click', searchWeather);
    cityInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            searchWeather();
        }
    });

    // Focus on input field when page loads
    cityInput.focus();

    function searchWeather() {
        const city = cityInput.value.trim();
        
        if (!city) {
            showError('Please enter a city name');
            return;
        }

        showLoading();
        hideError();
        hideWeatherCard();

        fetch(`/api/weather?city=${encodeURIComponent(city)}`)
            .then(response => {
                if (!response.ok) {
                    return response.text().then(text => {
                        throw new Error(text || `HTTP error! status: ${response.status}`);
                    });
                }
                return response.json();
            })
            .then(data => {
                hideLoading();
                displayWeather(data);
            })
            .catch(error => {
                hideLoading();
                console.error('Error fetching weather:', error);
                showError(getErrorMessage(error.message));
            });
    }

    function displayWeather(data) {
        cityName.textContent = data.city;
        countryName.textContent = data.country;
        temperature.textContent = Math.round(data.temperature);
        feelsLike.textContent = Math.round(data.feels_like);
        weatherDescription.textContent = data.description;
        humidity.textContent = data.humidity;
        pressure.textContent = data.pressure;
        windSpeed.textContent = data.wind_speed.toFixed(1);
        cloudCover.textContent = data.cloud_cover;
        
        // Set weather icon
        if (data.icon) {
            weatherIcon.src = `https://openweathermap.org/img/wn/${data.icon}@2x.png`;
            weatherIcon.alt = data.description;
        }

        showWeatherCard();
    }

    function showLoading() {
        loadingSpinner.classList.remove('hidden');
    }

    function hideLoading() {
        loadingSpinner.classList.add('hidden');
    }

    function showError(message) {
        errorText.textContent = message;
        errorMessage.classList.remove('hidden');
    }

    function hideError() {
        errorMessage.classList.add('hidden');
    }

    function showWeatherCard() {
        weatherCard.classList.remove('hidden');
    }

    function hideWeatherCard() {
        weatherCard.classList.add('hidden');
    }

    function getErrorMessage(errorMsg) {
        if (errorMsg.includes('404')) {
            return 'City not found. Please check the city name and try again.';
        } else if (errorMsg.includes('401')) {
            return 'API key error. Please check the server configuration.';
        } else if (errorMsg.includes('API key not configured')) {
            return 'Server configuration error. API key is missing.';
        } else if (errorMsg.includes('Failed to fetch') || errorMsg.includes('Network')) {
            return 'Network error. Please check your internet connection and try again.';
        } else {
            return 'Unable to fetch weather data. Please try again later.';
        }
    }
});