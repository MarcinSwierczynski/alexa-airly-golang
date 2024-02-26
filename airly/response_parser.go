package airly

import (
	"encoding/json"
	"log"
)

type response struct {
	Current current `json:"current"`
}

type current struct {
	Indexes   []index    `json:"indexes"`
	Standards []standard `json:"standards"`
}

type index struct {
	Description string  `json:"description"`
	Value       float32 `json:"value"`
}

type standard struct {
	Pollutant string  `json:"pollutant"`
	Percent   float32 `json:"percent"`
}

func parseJsonResponse(data []byte) response {
	var response response
	parsingError := json.Unmarshal(data, &response)
	if parsingError != nil {
		log.Fatalln(parsingError)
	}

	return response
}
