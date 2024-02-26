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

	if len(json.Current.Indexes[0].Description) <= 0 {
		t.Fatal("Description cannot be empty")
	}

	if json.Current.Indexes[0].Value <= 0 {
		t.Fatal("CAQI cannot be negative")
	}
}
