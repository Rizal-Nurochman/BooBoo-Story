package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"

	"github.com/BooBooStory/config"
)

type Service interface {
	SendPasswordResetEmail(toEmail, token string) error
}

type service struct {
		smtpHost     string
		smtpPort     string
		smtpUser     string
		smtpPassword string
		senderEmail  string
}

func NewService() Service {
	return &service{
		smtpHost:     config.Envs.SmtpHost,
		smtpPort:     config.Envs.SmtpPort,
		smtpUser:     config.Envs.SmtpUser,
		smtpPassword: config.Envs.SmtpPassword,
		senderEmail:  config.Envs.SmtpSenderEmail,
	}
}

func (s *service) SendPasswordResetEmail(toEmail, token string) error {
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
	subject := "Reset Password Akun BooBooStory Anda"

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", config.Envs.FE_URL, token)

	data := struct {
		ResetLink string
	}{
		ResetLink: resetLink,
	}

	body, err := parseEmailTemplate("reset_password_template.html", data)
	if err != nil {
		return fmt.Errorf("gagal mem-parse template email reset: %w", err)
	}

	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		body)

	auth := smtp.PlainAuth("", s.smtpUser, s.smtpPassword, s.smtpHost)

	err = smtp.SendMail(addr, auth, s.senderEmail, []string{toEmail}, msg)
	if err != nil {
		return fmt.Errorf("gagal mengirim email reset: %w", err)
	}

	return nil
}



func parseEmailTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("v1/email/templates/%s", templateFileName))
	if err != nil {
		return "", err
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
