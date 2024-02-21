package airly

import (
	"log"
	"net/http"
)

const endpoint = "https://airapi.airly.eu/v2/measurements/nearest"
const apiKey = "<api_key>"
const maxDistanceKm = "1"
const pm25pollutant = "PM25"

func fetchAirly(latitude, longitude string) *http.Response {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", endpoint, nil)
	request.Header.Set("apikey", apiKey)

	query := request.URL.Query()
	query.Add("lat", latitude)
	query.Add("lng", longitude)
	query.Add("maxDistanceKM", maxDistanceKm)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}
