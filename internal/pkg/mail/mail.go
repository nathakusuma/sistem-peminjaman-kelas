package mail

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"sync"

	"gopkg.in/gomail.v2"

	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/domain/mailtmpl"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/infrastructure/config"
	"github.com/nathakusuma/sistem-peminjaman-kelas/internal/pkg/log"
)

type IMailer interface {
	Send(recipientEmail, subject, templateName string, data map[string]any) error
}

type mailer struct {
	dialer    *gomail.Dialer
	templates *template.Template
}

var (
	instance IMailer
	once     sync.Once
)

func NewMailDialer() IMailer {
	once.Do(func() {
		// Parse all templates at startup
		templates, err := template.ParseFS(mailtmpl.Templates, "*.html")
		if err != nil {
			log.Fatal(context.Background()).Err(err).Msg("failed to parse email templates")
			return
		}

		instance = &mailer{
			dialer: gomail.NewDialer(
				config.GetEnv().SMTPHost,
				config.GetEnv().SMTPPort,
				config.GetEnv().SMTPUsername,
				config.GetEnv().SMTPPassword,
			),
			templates: templates,
		}
	})

	return instance
}

func (m *mailer) Send(recipientEmail, subject, templateName string, data map[string]any) error {
	var tmplOutput bytes.Buffer

	err := m.templates.ExecuteTemplate(&tmplOutput, templateName, data)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", "Sistem Peminjaman Kelas <"+config.GetEnv().SMTPEmail+">")
	mail.SetHeader("To", recipientEmail)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", tmplOutput.String())

	return m.dialer.DialAndSend(mail)
}
