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

// SendRuleViolationAlert sends an alert when a rule is violated
func (s *EmailService) SendRuleViolationAlert(ctx context.Context, fleetEmail string, vehicleModel string, message string, priority string) error {
	subject := fmt.Sprintf("[%s Priority] Vehicle Rule Violation Alert", priority)
	body := fmt.Sprintf(
		"Vehicle Rule Violation Detected\n\n"+
			"Vehicle: %s\n"+
			"Alert: %s\n"+
			"Please review the telemetry data and take necessary action.\n\n"+
			"This is an automated alert from MotorQ Fleet Management System.",
		vehicleModel,
		message,
	)

	return s.SendAlertEmail(ctx, fleetEmail, subject, body)
}