package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const QueueName = "alert_queue"

type Message struct {
	FleetEmail   string `json:"fleet_email"`
	VehicleModel string `json:"vehicle_model"`
	Message      string `json:"message"`
	Priority     string `json:"priority"`
	Timestamp    string `json:"timestamp"`
}

type Rabbitmq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitmq(url string) (*Rabbitmq, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	queue, err := channel.QueueDeclare(
		QueueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, err
	}

	return &Rabbitmq{
		conn:    conn,
		channel: channel,
		queue:   queue,
	}, nil
}

func (r *Rabbitmq) PublishAlert(ctx context.Context, msg Message) error {
	msg.Timestamp = time.Now().Format(time.RFC3339)

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = r.channel.PublishWithContext(ctx,
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		return err
	}

	log.Printf("📬 [RabbitMQ] Published %s priority alert: %s", msg.Priority, msg.Message)
	return nil
}

func (r *Rabbitmq) ConsumeAlerts(
	ctx context.Context,
	emailService EmailServiceInterface,
) error {
	err := r.channel.Qos(1, 0, false)
	if err != nil {
		return err
	}

	msgs, err := r.channel.Consume(
		r.queue.Name,
		"",
		false, // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Error consuming alerts: %v", err)
		return err
	}

	log.Println("📧 [RabbitMQ] Alert consumer started")

	for {
		select {
		case <-ctx.Done():
			log.Println("📧 [RabbitMQ] Alert consumer stopped")
			return ctx.Err()
		case msg, ok := <-msgs:
			if !ok {
				return nil
			}

			var alertMsg Message
			if err := json.Unmarshal(msg.Body, &alertMsg); err != nil {
				msg.Nack(false, false)
				log.Printf("Error unmarshalling alert: %v", err)
				continue
			}

			subject := fmt.Sprintf("[%s Priority] Vehicle Rule Violation Alert", alertMsg.Priority)
			body := fmt.Sprintf(
				"Vehicle Rule Violation Detected\n\n"+
					"Vehicle: %s\n"+
					"Alert: %s\n"+
					"Please review and take action.",
				alertMsg.VehicleModel,
				alertMsg.Message,
			)

			log.Printf("📧 [RabbitMQ] Processing %s alert for %s...", alertMsg.Priority, alertMsg.FleetEmail)

			// 2 second email delay
			time.Sleep(2 * time.Second)

			if err := emailService.SendAlertEmail(ctx, alertMsg.FleetEmail, subject, body); err != nil {
				msg.Nack(false, true) // requeue on failure
				log.Printf("Error sending email: %v", err)
				continue
			}

			msg.Ack(false)
			log.Printf("✅ [RabbitMQ] Alert processed: %s", alertMsg.Message)
		}
	}
}

func (r *Rabbitmq) GetQueueStats() (high, mid, low int) {
	// For simplicity, just return total message count
	info, err := r.channel.QueueInspect(r.queue.Name)
	if err != nil {
		return 0, 0, 0
	}
	return info.Messages, 0, 0 // simplified - just total
}

func (r *Rabbitmq) Close() error {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	return nil
}

// RabbitMQProcessor wraps Rabbitmq for alert processing
type RabbitMQProcessor struct {
	queue        *Rabbitmq
	emailService EmailServiceInterface
}

func NewRabbitMQProcessor(emailService EmailServiceInterface) (*RabbitMQProcessor, error) {
	url := "amqp://guest:guest@localhost:5672/"
	mq, err := NewRabbitmq(url)
	if err != nil {
		return nil, err
	}

	return &RabbitMQProcessor{
		queue:        mq,
		emailService: emailService,
	}, nil
}

func (p *RabbitMQProcessor) PublishAlert(msg Message) error {
	return p.queue.PublishAlert(context.Background(), msg)
}

func (p *RabbitMQProcessor) StartWorker(ctx context.Context) {
	p.queue.ConsumeAlerts(ctx, p.emailService)
}

func (p *RabbitMQProcessor) Stop() {
	// nothing to do, context cancellation stops the consumer
}

func (p *RabbitMQProcessor) GetQueueStats() (high, mid, low int) {
	return p.queue.GetQueueStats()
}

func (p *RabbitMQProcessor) Close() error {
	return p.queue.Close()
}

// Processor interface for queue operations
type Processor interface {
	PublishAlert(msg Message) error
	StartWorker(ctx context.Context)
	Stop()
	GetQueueStats() (high, mid, low int)
	Close() error
}

// EmailServiceInterface for email sending
type EmailServiceInterface interface {
	SendAlertEmail(ctx context.Context, to string, subject string, body string) error
}