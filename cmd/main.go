package main

import (
	"fmt"
	"github.com/emersion/go-smtp"
	"log"
	"smtpTest/internal/GetLetters"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	be := &GetLetters.Backend{}
	s := smtp.NewServer(be)
	s.Addr = ":25"
	s.Domain = "mx.mailer.ru"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	fmt.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}