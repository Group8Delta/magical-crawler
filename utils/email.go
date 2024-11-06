package utils

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	Host        string
	Port        string
	SenderEmail string
	SenderPass  string
}

func NewEmailService() *EmailService {
	return &EmailService{
		Host:        os.Getenv("EMAIL_HOST"),
		Port:        os.Getenv("EMAIL_PORT"),
		SenderEmail: os.Getenv("EMAIL_SENDER"),
		SenderPass:  os.Getenv("EMAIL_PASSWORD"),
	}
}

func (e *EmailService) SendEmail(recipientEmail string, recipientName string, file string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.SenderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", os.Getenv("EMAIL_SUBJECT"))
	m.Attach(file)

	body, err := loadTemplate("templates/email_template.html", recipientName)
	if err != nil {
		log.Printf("Failed to load email template: %v", err)
	}

	m.SetBody("text/html", body)

	port, err := strconv.Atoi(e.Port)
	if err != nil {
		log.Printf("Invalid port number: %v", err)
	}

	dailer := gomail.NewDialer(e.Host, port, e.SenderEmail, e.SenderPass)

	if err := dailer.DialAndSend(m); err != nil {
		log.Printf("Could not send email to %v: %v", recipientEmail, err) // assign error to the admin
		return err
	}
	log.Printf("Email sent successfully to %v", recipientEmail)
	return nil
}

func loadTemplate(templateFileName string, recipientName string) (string, error) {
	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	var template bytes.Buffer
	if err := tmpl.Execute(&template, recipientName); err != nil {
		return "", err
	}
	return template.String(), nil
}
