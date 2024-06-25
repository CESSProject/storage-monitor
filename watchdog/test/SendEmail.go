package test

import (
	"bytes"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/model"
	"gopkg.in/gomail.v2"
	"html/template"
	"time"
)

func SendMail() error {
	content := model.AlertContent{
		AlertTime:     time.Now().Format(constant.TimeFormat),
		HostIp:        "127.0.0.1",
		ContainerName: "miner1",
		Description:   "The Storage Miner is not a positive status or get punishment",
	}

	tmpl, err := template.ParseFiles("./internal/util/template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, content); err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "xxxxx@cess.cloud")
	m.SetHeader("To", "xxxxx@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.example.com", 80, "xxxxx@cess.cloud", "xxxxx")
	if err = d.DialAndSend(m); err != nil {
		return err
	}
	return nil

}
