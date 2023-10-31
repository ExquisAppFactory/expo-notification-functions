package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
)

// SendSlackNotification sends a Slack notification to the specified webhook URL
func SendSlackNotification(message string, webhookURL string) {
	log.Printf("Sending slack notification: %s", []byte(message))

	slackMessage := []byte(`{"text":"` + message + `"}`)

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(slackMessage))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-200 response: %d", resp.StatusCode)
	}

	log.Println("Slack notification sent!")
}

// safeCompare compares two strings in constant time
func safeCompare(a string, b string) bool {
	return hmac.Equal([]byte(a), []byte(b))
}

// VerifySignature verifies the signature of the request body
func VerifySignature(signature string, body []byte, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)

	return safeCompare(signature, fmt.Sprintf("sha256=%x", expectedMAC))
}
