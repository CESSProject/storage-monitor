package util

import (
	"bytes"
	"fmt"
	"github.com/CESSProject/watchdog/internal/model"
	"log"
	"net/smtp"
	"text/template"
)

type SmtpConfig struct {
	SmtpUrl      string
	SmtpPort     string
	SenderAddr   string
	SmtpPassword string
	Receiver     []string
}

func (m SmtpConfig) SendMail(content model.MailContent) (err error) {
	auth := smtp.PlainAuth("", m.SenderAddr, m.SmtpPassword, m.SmtpUrl)
	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		return
	}
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Miner Status Alert\n%s\n\n", mimeHeaders)))
	err = t.Execute(&body, content)
	if err != nil {
		log.Println("Error executing template:", err)
		return
	}
	err = smtp.SendMail(m.SmtpUrl+":"+m.SmtpPort, auth, m.SenderAddr, m.Receiver, body.Bytes())
	if err != nil {
		log.Println("Error sending email:", err)
		return
	}
	return
}
