package notification

import (
	"context"
	"fmt"
	"time"
)

type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

// SendAlertEmail simulates sending an email (takes 2 seconds)
func (s *EmailService) SendAlertEmail(ctx context.Context, to string, subject string, body string) error {
	// Simulate email sending delay (2 seconds)
	time.Sleep(2 * time.Second)

	// Log the email details
	fmt.Printf("📧 [MOCK EMAIL SENT]\n")
	fmt.Printf("   To: %s\n", to)
	fmt.Printf("   Subject: %s\n", subject)
	fmt.Printf("   Body: %s\n", body)
	fmt.Printf("   Status: Delivered (mock)\n")
	fmt.Printf("   Time: %s\n\n", time.Now().Format(time.RFC3339))

	return nil
}