package main

import "os"

func sendEmail(a string) {
	senderEmail = os.Getenv("EMAIL")

}
