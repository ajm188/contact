package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

var (
	hostname = os.Getenv("CONTACT_HOSTNAME")
	username = os.Getenv("CONTACT_USERNAME")
	password = os.Getenv("CONTACT_PASSWORD")
)

var (
	contactAddr = fmt.Sprintf("contact@%s", hostname)
)

var (
	auth = smtp.PlainAuth("", username, password, hostname)
)

func handleContact(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// TODO: 400
		return
	}

	replyTo := r.FormValue("replyTo")
	subject := r.FormValue("subject")
	body := r.FormValue("body")

	if replyTo == "" || subject == "" || body == "" {
		// TODO: 400
		return
	}

	// TODO: sanitize user input

	message := fmt.Sprintf("From: %s\r\nReply-To: %s\r\nSubject: %s\r\n\r\n%s\r\n.\r\n", replyTo, replyTo, subject, body)
	msg := []byte(message)
	if err := smtp.SendMail(hostname+":25", auth, contactAddr, []string{contactAddr}, msg); err != nil {
		// TODO: 500
		return
	}
	// TODO: 200
}

func main() {
	// TODO: make sure hostname, username, password are non-empty
	// TODO: make sure port is non-empty
	http.HandleFunc("/contact", handleContact)
	http.ListenAndServe(os.Getenv("CONTACT_PORT"), nil)
}
