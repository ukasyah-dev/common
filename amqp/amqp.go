package amqp

import (
	"context"

	"github.com/bytedance/sonic"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

var conn *amqp091.Connection
var ch *amqp091.Channel

func Open(amqpURL string) {
	var err error

	conn, err = amqp091.Dial(amqpURL)
	if err != nil {
		panic(err)
	}

	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}
}

func Close() error {
	ch.Close()
	conn.Close()
	return nil
}

func DeclareQueues(names ...string) error {
	for _, name := range names {
		_, err := ch.QueueDeclare(name, true, false, false, false, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func Publish(ctx context.Context, queue string, data any) error {
	body, err := sonic.Marshal(data)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, "", queue, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}

func Consume[I, O any](
	ctx context.Context,
	queue string,
	consumerGroupID string,
	f func(context.Context, *I) (*O, error),
) error {
	msgs, err := ch.Consume(
		queue,           // queue
		consumerGroupID, // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg := <-msgs
			data := new(I)

			err := sonic.Unmarshal(msg.Body, data)
			if err != nil {
				return err
			}

			_, err = f(ctx, data)
			if err != nil {
				return err
			}
		}
	}
}
