package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func safeCompare(a, b string) bool {
	return hmac.Equal([]byte(a), []byte(b))
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	expoSignature := r.Header.Get("expo-signature")

	log.Printf("expo-signature from Header: %s", expoSignature)

	secretWebhookKey := os.Getenv("SECRET_WEBHOOK_KEY")

	h := hmac.New(sha1.New, []byte(secretWebhookKey))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	h.Write(body)
	log.Printf("expo-signature from Body: %s", hex.EncodeToString(h.Sum(nil)))
	hash := "sha1=" + hex.EncodeToString(h.Sum(nil))

	if !safeCompare(expoSignature, hash) {
		http.Error(w, "Signatures didn't match!", http.StatusInternalServerError)
		return
	}
}

func sendSlackNotifications(message string) {
	url := os.Getenv("SLACK_WEBHOOK_URL")

	buildMessage := "New Build Artifact for For Android Generated! 🚀🚀🚀\n Click to install " + message

	data := []byte(`{"text":"` + buildMessage + `"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	log.Printf("request: %v", req)

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
