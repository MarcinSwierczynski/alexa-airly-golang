package airly

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

func parseJsonResponse(resp *http.Response) response {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response response
	parsingError := json.Unmarshal(data, &response)
	if parsingError != nil {
		log.Fatalln(parsingError)
	}

	return response
}
