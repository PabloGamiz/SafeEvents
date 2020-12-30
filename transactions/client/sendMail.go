package client

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"
)

// txAddFav represents an
type txSendMail struct {
	request []string
	body    string
	ctx     context.Context
}

func (tx *txSendMail) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	tx.body = buf.String()
	return nil
}

func (tx *txSendMail) SendEmail() error {
	username := os.Getenv(envMailUsr)
	password := os.Getenv(envMailPwd)
	emailServer := os.Getenv(envSMTPHost)
	addr := emailServer + ":" + os.Getenv(envSMTPPort)

	log.Printf("\n" + username + "\n" + password + "\n" + emailServer + "\n" + addr)

	auth := smtp.PlainAuth("",
		username,
		password,
		emailServer)

	mime := emailMIME
	subject := emailSubject
	msg := []byte(subject + mime + "\n" + tx.body)

	err := smtp.SendMail(addr, auth, username, tx.request, msg)

	if err != nil {
		return fmt.Errorf("ERROR: attempting to send a mail %v", err)
	}

	return nil
}

// Precondition validates the transaction is ready to run
func (tx *txSendMail) Precondition() (err error) {
	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txSendMail) Postcondition(ctx context.Context) (v interface{}, err error) {

	templateData := struct {
		URL string
	}{
		URL: emailInfoURL,
	}

	if err = tx.ParseTemplate("template.html", templateData); err == nil {
		err = tx.SendEmail()
	}

	return
}

// Commit commits the transaction result
func (tx *txSendMail) Commit() error {
	return nil
}

// Rollback rollbacks any change caused while the transaction
func (tx *txSendMail) Rollback() {

}
