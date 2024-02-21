package airly

import (
	"log"
)

type AirQuality struct {
	Description string
	Caqi        float32
	Pm25Percent float32
}

func GetAirQuality() AirQuality {
	response := fetchAirly("50.273966", "19.004781")
	parsedJson := parseJsonResponse(response)
	standard, _ := findByPollutant(parsedJson.Current.Standards, pm25pollutant)

	airQuality := AirQuality{
		Description: parsedJson.Current.Indexes[0].Description,
		Caqi:        parsedJson.Current.Indexes[0].Value,
		Pm25Percent: standard.Percent,
	}

	log.Println(airQuality)
	return airQuality
}

func findByPollutant(standards []standard, pollutant string) (standard, bool) {
	for _, standard := range standards {
		if standard.Pollutant == pollutant {
			return standard, true
		}
	}
	return standard{Percent: -1}, false
}
