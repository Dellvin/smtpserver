package GetLetters

import (
	"fmt"
	"github.com/emersion/go-smtp"
	"io"
	"io/ioutil"
	send "smtpTest/internal/SendLetters"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	fmt.Println("USPEX")
	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	fmt.Println("USPEX Anonymous ", state.Hostname, state.RemoteAddr, state.LocalAddr)
	fmt.Println(state.TLS)

	return &Session{}, nil
}

// A Session is returned after successful login.
type Session struct{}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	fmt.Println("HUI")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in mail", r)
		}
	}()
	fmt.Println("EMail from:", from, opts.Auth)
	go send.SendAnswer2(from)


	return nil
}

func (s *Session) Rcpt(to string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in rcpt", r)
		}
	}()
	fmt.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		fmt.Println("Data:", string(b))
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}