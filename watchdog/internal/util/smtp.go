package util

import (
	"bytes"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"gopkg.in/gomail.v2"
	"text/template"
)

type SmtpConfig struct {
	SmtpUrl      string
	SmtpPort     int
	SenderAddr   string
	SmtpPassword string
	Receiver     []string
}

func (conf *SmtpConfig) SendMail(content model.AlertContent) (err error) {

	t, err := template.ParseFiles("./internal/util/template.html")
	if err != nil {
		log.Logger.Errorf("Error parsing template: %v", err)
		return err
	}
	var body bytes.Buffer
	err = t.Execute(&body, content)
	if err != nil {
		log.Logger.Errorf("Error executing template: %v", err)
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", conf.SenderAddr)

	m.SetHeader("Subject: Storage Miner Status Alert!")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(conf.SmtpUrl, 80, conf.SenderAddr, conf.SmtpPassword)
	for i := 0; i < constant.HttpMaxRetry; i++ {
		for _, receiver := range conf.Receiver {
			m.SetHeader("To", receiver)
			if err = d.DialAndSend(m); err != nil {
				log.Logger.Warnf("Fail when send email: %v, retrying (%d/%d)\n", err, i+1, constant.HttpMaxRetry)
			} else {
				break
			}
		}
	}
	return
}
