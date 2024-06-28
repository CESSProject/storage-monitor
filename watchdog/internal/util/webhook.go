package util

import (
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"net/http"
	"strings"
)

type WebhookConfig struct {
	Webhooks []string
}

func (conf *WebhookConfig) SendAlertToWebhook(content model.AlertContent) (err error) {
	sendData := `{
		"msg_type": "text",
		"content": {"text": "` + "CESS Information: " + " Alert Time: " + content.AlertTime + ", Host: " + content.HostIp + ", Miner Name: " + content.ContainerName + ", Description: " + content.Description + `"}
	}`
	for i := 0; i < len(conf.Webhooks); i++ {
		req, err := http.NewRequest("POST", conf.Webhooks[i], strings.NewReader(sendData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Logger.Error("Failed to create new http request client when call a webhook")
			return err
		}
		go func() {
			err := sendRequest(req)
			if err != nil {
				return
			}
		}()
	}
	return nil
}

func sendRequest(req *http.Request) error {
	client := &http.Client{}
	for j := 0; j < constant.HttpMaxRetry; j++ {
		resp, err := client.Do(req)
		if err != nil {
			log.Logger.Warnf("Fail when request to webhook: %v, retrying (%d/%d)\n", err, j+1, constant.HttpMaxRetry)
		} else {
			if resp.StatusCode != http.StatusOK {
				log.Logger.Warnf("Unexceptional response status code: %d", resp.StatusCode)
			}
			err := resp.Body.Close()
			if err != nil {
				return err
			}
			break
		}
		err = resp.Body.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
