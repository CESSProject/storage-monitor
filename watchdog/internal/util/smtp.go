package util

import (
	"bytes"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/model"
	"log"
	"net/smtp"
	"strconv"
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
	auth := smtp.PlainAuth("", conf.SenderAddr, conf.SmtpPassword, conf.SmtpUrl)
	t, err := template.ParseFiles("./internal/util/template.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		return err
	}
	var body bytes.Buffer
	err = t.Execute(&body, content)
	if err != nil {
		log.Println("Error executing template:", err)
		return err
	}
	subject := "Subject: Storage Miner Status Alert!"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + body.String())
	for i := 0; i < constant.HttpMaxRetry; i++ {
		err = smtp.SendMail(conf.SmtpUrl+":"+strconv.Itoa(conf.SmtpPort), auth, conf.SenderAddr, conf.Receiver, msg)
		if err != nil {
			log.Printf("Fail when send email: %v, retrying (%d/%d)\n", err, i+1, constant.HttpMaxRetry)
		} else {
			break
		}
	}
	return
}
