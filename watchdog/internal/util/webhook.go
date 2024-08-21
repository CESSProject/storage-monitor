package util

import (
	"bytes"
	"encoding/json"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"net/http"
	"sync"
)

type WebhookSender interface {
	SendMessage(message string) error
}

type DiscordWebhook struct {
	// https://discord.com/api/webhooks/................
	WebhookURL string
}

func (discord *DiscordWebhook) SendMessage(message string) error {
	payload := map[string]interface{}{
		"content": message,
	}
	return sendWebhookRequest(discord.WebhookURL, payload)
}

type TeamsWebhook struct {
	// api.telegram.org/................
	WebhookURL string
}

func (teams *TeamsWebhook) SendMessage(message string) error {
	payload := map[string]interface{}{
		"text": message,
	}
	return sendWebhookRequest(teams.WebhookURL, payload)
}

type WechatWebhook struct {
	// qyapi.weixin.qq.com/................
	WebhookURL string
}

func (wechat *WechatWebhook) SendMessage(message string) error {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": message,
		},
	}
	return sendWebhookRequest(wechat.WebhookURL, payload)
}

type SlackWebhook struct {
	// https://hooks.slack.com/services/................
	WebhookURL string
}

func (slack *SlackWebhook) SendMessage(message string) error {
	payload := map[string]interface{}{
		"text": message,
	}
	return sendWebhookRequest(slack.WebhookURL, payload)
}

type DingTalkWebhook struct {
	// https://oapi.dingtalk.com/robot/send?access_token=................
	WebhookURL string
}

func (ding *DingTalkWebhook) SendMessage(message string) error {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": message,
		},
	}
	return sendWebhookRequest(ding.WebhookURL, payload)
}

type LarkWebhook struct {
	// https://open.larksuite.com/open-apis/bot/v2/hook/...............
	WebhookURL string
}

func (lark *LarkWebhook) SendMessage(message string) error {
	payload := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": message,
		},
	}
	return sendWebhookRequest(lark.WebhookURL, payload)
}

func sendWebhookRequest(url string, payload interface{}) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	for j := 0; j < constant.HttpMaxRetry; j++ {
		resp, err := http.Post(url, constant.HttpPostContentType, bytes.NewBuffer(jsonPayload))
		if err != nil {
			log.Logger.Warnf("Fail when request to webhook: %v, retrying (%d/%d)\n", err, j+1, constant.HttpMaxRetry)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			log.Logger.Warnf("Unexceptional response status code: %d", resp.StatusCode)
		}
		err = resp.Body.Close()
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusOK {
			break
		}
	}
	return nil
}

type WebhookConfig struct {
	Webhooks []string
}

func (conf *WebhookConfig) SendAlertToWebhook(content model.AlertContent) (err error) {
	var wg sync.WaitGroup
	errChan := make(chan error, len(conf.Webhooks))
	message := "CESS Information: " + " Alert Time: " + content.AlertTime + ", Host: " + content.HostIp + ", Miner Name: " + content.ContainerName + ", Description: " + content.Description
	for _, url := range conf.Webhooks {
		var hook WebhookSender
		switch GetWebhookType(url) {
		case "discord":
			hook = &DiscordWebhook{url}
			break
		case "slack":
			hook = &SlackWebhook{url}
			break
		case "teams":
			hook = &TeamsWebhook{url}
			break
		case "lark":
			hook = &LarkWebhook{url}
			break
		case "ding":
			hook = &DingTalkWebhook{url}
			break
		case "wechat":
			hook = &WechatWebhook{url}
			break
		default:
			log.Logger.Warn("Unknown webhook type, cannot send webhook alert")
			return nil
		}
		wg.Add(1)
		go func(h WebhookSender) {
			defer wg.Done()
			if err := h.SendMessage(message); err != nil {
				errChan <- err
			}
		}(hook)
	}
	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			log.Logger.Errorf("Error sending webhook: %v", err)
		}
	}
	return nil
}
