package main

import (
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"crypto/tls"
)

func sendAnswer(email string){
	from := mail.Address{"", "test@mailer.ru.com"}
	to   := mail.Address{"", email}
	subj := "Hello"
	body := "We are happy to see you in our alfa smtp-test!"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "localhost:1025"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("","danya.tchyorny@yandex.ru", "CherDan9851681553", host)

	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
	fmt.Println("Sent answer to: ", email)
}