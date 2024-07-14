package mail

import (
	"net/url"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var smtpClient *mail.SMTPClient

func Open(smtpURL string) {
	parsed, err := url.Parse(smtpURL)
	if err != nil {
		panic(err)
	}

	client := mail.NewSMTPClient()

	if parsed.Hostname() != "" {
		client.Host = parsed.Hostname()
	}

	if parsed.Port() != "" {
		client.Port, _ = strconv.Atoi(parsed.Port())
	}

	if parsed.User.Username() != "" {
		client.Username = parsed.User.Username()
	}

	if _, ok := parsed.User.Password(); ok {
		client.Password, _ = parsed.User.Password()
	}

	if parsed.Scheme == "smtps" {
		client.Encryption = mail.EncryptionSTARTTLS
	}

	client.KeepAlive = true

	smtpClient, err = client.Connect()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			smtpClient.Noop()
			time.Sleep(30 * time.Second)
		}
	}()
}

func Close() error {
	return smtpClient.Close()
}

func Send(e Email) error {
	email := mail.NewMSG()
	email.SetFrom(e.From)
	email.AddTo(e.To)
	email.SetSubject(e.Subject)
	email.SetBody(mail.TextHTML, e.Body.GetHTML())

	return email.Send(smtpClient)
}
