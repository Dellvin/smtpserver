package main

import (
	"crypto/tls"
	"fmt"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"log"
	"net"
	"net/mail"
	baseSMTP "net/smtp"
	"strings"
)

func sendAnswer(email string){
		if email=="bot@mailer.ru.com"{
			return
		}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer", r)
		}
	}()
	fmt.Println("HUI_1")
	from := mail.Address{"", "bot@mailer.ru.com"}
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
	servername := "localhost:25"

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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in getanswer2", r)
		}
	}()
	fmt.Println("KEK_1")
	// Set up authentication information.
	if email=="bot@mailer.ru.com"{
		return
	}
	auth := sasl.NewPlainClient("", "bot@mailer.ru.com", "password")
	fmt.Println("KEK_2")
	servername := "mailer.ru.com:25"
	to := []string{email}
	msg := strings.NewReader("To: "+email+"\r\n" +
		"Subject: Hello SMTP!!!\r\n" +
		"\r\n" +
		"We are happy to see you in our alfa smtp-test!\r\n")
	fmt.Println("KEK_3")
	err := smtp.SendMail(servername, auth, "bot@mailer.ru.com", to, msg)
	fmt.Println("KEK_4")
	if err != nil {
		fmt.Println("Error in sendAnswer2", err.Error())
	}
	fmt.Println("success sendAnswer2")
}