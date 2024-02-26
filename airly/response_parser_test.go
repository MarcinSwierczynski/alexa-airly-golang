package airly

import (
	"log"
	"os"
	"testing"
)

func TestParseJsonResponse(t *testing.T) {
	file, err := os.ReadFile("resources/api-response.json")
	if err != nil {
		log.Fatalln(err)
	}

	json := parseJsonResponse(file)

	if json.Current.Indexes[0].Description != "Great air here today!" {
		t.Fatal("Description needs to be proper")
	}

	if json.Current.Indexes[0].Value != 7.43 {
		t.Fatal("CAQI needs to be proper")
	}

	if json.Current.Standards[0].Percent != 29.71 {
		t.Fatal("PM2.5 needs to be proper")
	}
}
