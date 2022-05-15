package services

import (
	"bytes"
	"fmt"
	"github.com/Zaida-3dO/goblin/config"
	"github.com/Zaida-3dO/goblin/pkg/errs"
	mailgun "github.com/mailgun/mailgun-go"
	"html/template"
)

type EmailServiceInterface interface {
	SendForgotPasswordEmail(name, email, passwordResetToken, redirectTo string) *errs.Err
	SendPasswordResetEmail(name, email string) *errs.Err
}

type emailService struct{}

func NewEmailService() EmailServiceInterface {
	return &emailService{}
}

type emailDetails struct {
	body         string
	recipient    string
	subject      string
	htmlTemplate string
}

func newEmailDetails(body, recipient, subject, htmlTemplate string) *emailDetails {
	return &emailDetails{
		body,
		recipient,
		subject,
		htmlTemplate,
	}
}

func (es *emailService) SendForgotPasswordEmail(name, email, passwordResetToken, redirectTo string) *errs.Err {
	//fmt.Printf("name: %s, email: %s, passwordResetToken: %s, redirectTo: %s\n",
	//	name, email, passwordResetToken, redirectTo)
	htmlTemplate, err := es.getHTMLAsString("forgotPassword.html", "./pkg/emails/forgotPassword.html", struct {
		Name         string
		RecoveryLink string
	}{
		name,
		fmt.Sprintf("%s/%s", redirectTo, passwordResetToken),
	})
	if err != nil {
		return err
	}

	details := newEmailDetails("", email, "Forgot Password", htmlTemplate)

	if err = es.sendEmail(details); err != nil {
		return err
	}

	return nil
}

func (es *emailService) SendPasswordResetEmail(name, email string) *errs.Err {
	htmlTemplate, err := es.getHTMLAsString("passwordChanged.html", "./pkg/emails/passwordChanged.html", struct {
		Name string
	}{
		name,
	})
	if err != nil {
		return err
	}

	details := newEmailDetails("", email, "Reset Password", htmlTemplate)

	if err = es.sendEmail(details); err != nil {
		return err
	}

	return nil
}

func (es *emailService) sendEmail(details *emailDetails) *errs.Err {
	mg, sender := mailgun.NewMailgun(config.Cfg.MgDomain, config.Cfg.MgAPIKey), config.Cfg.MgEmailTo

	message := mg.NewMessage(sender, details.subject, details.body, details.recipient)
	if details.htmlTemplate != "" {
		message.SetHtml(details.htmlTemplate)
	}

	_, _, err := mg.Send(message)
	if err != nil {
		return errs.NewInternalServerErr(err.Error(), err)
	}

	return nil
}

func (es *emailService) getHTMLAsString(name, path string, data interface{}) (string, *errs.Err) {
	temp, err := template.New(name).ParseFiles(path)
	if err != nil {
		return "", errs.NewInternalServerErr(err.Error(), err)
	}

	var buf bytes.Buffer
	err = temp.Execute(&buf, data)
	if err != nil {
		return "", errs.NewInternalServerErr(err.Error(), err)
	}

	return buf.String(), nil
}
