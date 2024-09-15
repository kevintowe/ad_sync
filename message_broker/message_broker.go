package message_broker

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Azure/go-amqp"
)

// MessageBrokerClient holds the AMQP connection instance
type MessageBrokerClient struct {
	Conn *amqp.Conn
}

// NewMessageBrokerClient initializes a new message broker connection with connection options
func NewMessageBrokerClient(ctx context.Context, amqpURL string) (*MessageBrokerClient, error) {
	// Establish a new connection with options (e.g., SASL authentication)
	conn, err := amqp.Dial(ctx, amqpURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to message broker: %v", err)
	}

	return &MessageBrokerClient{Conn: conn}, nil
}

// Close closes the message broker connection
func (mbc *MessageBrokerClient) Close() error {
	return mbc.Conn.Close()
}

// ConsumeFromTopic listens for messages from a topic and passes them to a handler function
func (mbc *MessageBrokerClient) ConsumeFromTopic(ctx context.Context, topicName, subscriptionName string, handler func(context.Context, *amqp.Message, *sync.WaitGroup)) error {
	// Open a session using the connection
	session, err := mbc.Conn.NewSession(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to create AMQP session: %v", err)
	}

	// Construct the subscription receiver address
	subscriptionPath := fmt.Sprintf("%s/Subscriptions/%s", topicName, subscriptionName)

	// Create a receiver for the subscription with the correct LinkSource and LinkCredit options
	receiver, err := session.NewReceiver(
		amqp.Link{Address: subscriptionPath}, // Set the source address
		amqp.LinkCredit(10),                  // Set prefetch credit
	)
	if err != nil {
		return fmt.Errorf("failed to create AMQP receiver: %v", err)
	}
	defer receiver.Close(ctx)

	// Receive messages in a loop
	for {
		msg, err := receiver.Receive(ctx, nil)
		if err != nil {
			log.Printf("failed to receive message: %v", err)
			break
		}

		// Use a WaitGroup to ensure the handler completes
		var wg sync.WaitGroup
		wg.Add(1)

		// Handle the message asynchronously
		go handler(ctx, msg, &wg)

		wg.Wait() // Ensure the handler completes before acknowledging the message

		// Acknowledge the message using receiver.AcceptMessage()
		if err := receiver.AcceptMessage(ctx, msg); err != nil {
			log.Printf("failed to acknowledge message: %v", err)
		}
	}

	return nil
}
