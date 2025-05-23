package service

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type EmailService struct {
	client *gmail.Service
}

func NewEmailService(ctx context.Context, credentialsFile string) (*EmailService, error) {
	data, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(data, gmail.GmailSendScope)
	if err != nil {
		return nil, err
	}

	client, err := gmail.NewService(ctx, option.WithHTTPClient(config.Client(ctx)))
	if err != nil {
		return nil, err
	}

	return &EmailService{client: client}, nil
}

func (s *EmailService) SendOrderConfirmation(ctx context.Context, to, orderID string) error {
	message := &gmail.Message{
		Raw: base64.URLEncoding.EncodeToString([]byte(
			"To: " + to + "\r\n" +
				"Subject: Order Confirmation\r\n" +
				"\r\n" +
				"Your order " + orderID + " has been confirmed!",
		)),
	}

	_, err := s.client.Users.Messages.Send("me", message).Do()
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	return nil
}
