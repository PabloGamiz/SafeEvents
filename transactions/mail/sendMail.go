package mail

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"os"
	"text/template"

	clientMOD "github.com/PabloGamiz/SafeEvents-Backend/model/client"
)

// txSendMail
type txSendMail struct {
	request    []uint
	recipients []string
	body       string
	ctx        context.Context
}

func (tx *txSendMail) ParseTemplate(templateFileName string) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, nil); err != nil {
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

	auth := smtp.PlainAuth("",
		username,
		password,
		emailServer)

	mime := emailMIME
	subject := emailSubject
	msg := []byte(subject + mime + "\n" + tx.body)

	err := smtp.SendMail(addr, auth, username, tx.recipients, msg)

	if err != nil {
		return fmt.Errorf("ERROR: attempting to send a mail %v", err)
	}

	return nil
}

// Precondition validates the transaction is ready to run
func (tx *txSendMail) Precondition() (err error) {
	tx.recipients = make([]string, len(tx.request))
	for index, clientID := range tx.request {
		if tx.recipients[index], err = clientMOD.FindClientEmailByClientID(tx.ctx, clientID); err != nil {
			return
		}
	}

	return nil
}

// Postcondition creates new user and a opens its first session
func (tx *txSendMail) Postcondition(ctx context.Context) (v interface{}, err error) {
	if err = tx.ParseTemplate("template/template.html"); err == nil {
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
