package main

import (
	"crypto/tls"
	"fmt"
	"github.com/emersion/go-sasl"
	"log"
	"net"
	"net/mail"
	baseSMTP "net/smtp"
	"github.com/emersion/go-smtp"
	"strings"
)

func sendAnswer(email string, opts smtp.MailOptions){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer", r)
		}
	}()
	fmt.Println("HUI_1")
	from := mail.Address{"", "test@mailer.ru.com"}
	to   := mail.Address{"", email}
	subj := "Hello"
	body := "We are happy to see you in our alfa smtp-test!"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	fmt.Println("HUI_2")
	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "mailer.ru:25"

	host, _, _ := net.SplitHostPort(servername)
	fmt.Println("HUI_3")
	auth := baseSMTP.PlainAuth("",from.String(), "keklol123", host)
	fmt.Println("HUI_4")
	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_5")
	c, err := baseSMTP.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_6")
	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_7")
	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_8")
	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_9")
	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_10")
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_11")
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("HUI_12")
	c.Quit()
	fmt.Println("Sent answer to: ", email)
}


func sendAnswer2(email string){
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "bot@mailer.ru.com", "password")
	servername := "localhost:1025"


	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{email}
	msg := strings.NewReader("To: "+email+"\r\n" +
		"Subject: Hello SMTP!!!\r\n" +
		"\r\n" +
		"We are happy to see you in our alfa smtp-test!\r\n")
	err := smtp.SendMail(servername, auth, "bot@mailer.ru.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}