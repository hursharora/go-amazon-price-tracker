package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func sendEmail(i item) {
	senderEmail := os.Getenv("GMAIL_ADDRESS")
	senderPassword := os.Getenv("GMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := fmt.Sprintf("The price of %s has dropped below your target price of $%.2f", i.name, i.targetPrice)

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	rec := []string{i.email}
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, rec, []byte(message))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Emailed %s", i.email)
}
