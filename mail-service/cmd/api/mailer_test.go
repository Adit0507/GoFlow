package main

import "testing"

func TestSendSMTPMessage(t *testing.T) {
	mailer := Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Username:    "",
		Password:    "",
		Encryption:  "none",
		FromAddress: "test@example.com",
		FromName:    "Test Sender",
	}

	msg := Message{
		To:      "receiver@example.com",
		Subject: "Test Email Subject",
		Data: map[string]any{
			"name":    "John Doe",
			"content": "This is a test email sent using Mailhog.",
		},
	}

	err := mailer.SendSMTPMessage(msg)
	if err != nil {
		t.Fatalf("expected no eror, got %v", err)
	}

}
