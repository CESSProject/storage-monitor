package util

import (
	"bytes"
	"encoding/json"
	"github.com/CESSProject/watchdog/internal/model"
	"log"
	"net/http"
)

func sendAlertToWebhook(webhookURL, message string) error {
	msg := model.WebhookContent{
		MsgType: "text",
	}
	msg.Content.Text = message
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("Encode webhook content failed")
		return err
	}
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Fatal("Failed to create new http request")
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed to send a http post request")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Unexceptional response status code: ", resp.StatusCode)
	}
	return nil
}
