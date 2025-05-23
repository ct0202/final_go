package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare("order_notifications", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func (r *RabbitMQ) PublishOrderNotification(ctx context.Context, orderID, email string) error {
	body, err := json.Marshal(map[string]string{"order_id": orderID, "email": email})
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"",
		"order_notifications",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	return nil
}

func (q *RabbitMQ) Close() error {
	return q.conn.Close()
}
