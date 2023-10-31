package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"

	"github.com/ExquisAppFactory/expo-notification-functions/pkg"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch {
	case request.HTTPMethod == "POST" && request.Resource == "/webhook":
		return EASBuildWebhook(request)
	case request.HTTPMethod == "GET" && request.Resource == "/webhook":
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "OK",
		}, nil
	case request.HTTPMethod == "GET" && request.Resource == "/":
		return InvocationData(request)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Page Not found",
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}

func EASBuildWebhook(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	expoSignature := req.Headers["expo-signature"]

	log.Printf("expo-signature from Header: %s", expoSignature)

	var data BuildDetails

	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error unmarshaling response",
		}, nil
	}

	if data.Platform != "android" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "OK",
		}, nil
	}

	if data.Platform == "android" && data.Status == "finished" {
		log.Println("Device: Android and Status: Finished detected")
		pkg.SendSlackNotification(data.Artifacts.BuildUrl)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}

func InvocationData(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request Body: %s", req.Body)

	meta := req.Headers["X-Invocation-Id"]

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       meta,
	}, nil
}

//func processEASBuildRequest(body string, secret string, webhookURL string) {
//	var data BuildDetails
//	err := json.Unmarshal(body, &data)
//	if err != nil {
//		fmt.Println("Error unmarshaling response:", err)
//		return
//	}
//
//	if err != nil {
//		fmt.Println("Error unmarshaling response:", err)
//		return
//	}
//
//	if data.Platform != "android" {
//		return
//	}
//
//	if data.Platform == "android" && data.Status == "finished" {
//		sendSlackNotification(data.Artifacts.BuildUrl)
//	}
//
//	w.Write([]byte("OK!"))
//}
