package main

import (
	"alexa-airly-golang/airly"
	"context"
	"fmt"
	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request alexa.Request) (alexa.Response, error) {
	airQuality := airly.GetAirQuality()

	message := fmt.Sprintf("%s The Common Air Quality Index is %d. The PM2.5 is %d%%",
		airQuality.Description, int(airQuality.Caqi), int(airQuality.Pm25Percent))

	return alexa.NewSimpleResponse("AlexaAirly", message), nil
}

func main() {
	lambda.Start(HandleRequest)
}
