package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func sendEmail(i item) {
	senderEmail := os.Getenv("GMAIL_ADDRESS")
	senderPassword := os.Getenv("GMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	to := fmt.Sprintf("To: %s", i.email)
	subject := "Subject: Go Amazon Price Tracker - Price Drop"
	message := fmt.Sprintf("The price of %s has dropped below your target price of $%.2f", i.name, i.targetPrice)

	fullMessage := strings.Join([]string{to, subject, message}, "\r\n")

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	rec := []string{i.email}
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, rec, []byte(fullMessage))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Emailed %s", i.email)
}
