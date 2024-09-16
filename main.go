package main

import (
	"context"
	"log"

	"github.com/Azure/go-amqp"
)

func main() {
	ctx := context.TODO()

	// create connection
	conn, err := amqp.Dial(ctx, "amqp://localhost", nil)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer conn.Close()

	// // open a session
	// session, err := conn.NewSession(ctx, nil)
	// if err != nil {
	// 	log.Fatal("Creating AMQP session:", err)
	// }

	// session.Close(ctx)
	// log.Printf("ending")
	// cfg := load_config.LoadConfig()

	// // define the topics and their corresponding handlers
	// topics := []struct {
	// 	topicName        string
	// 	subscriptionName string
	// 	handler          func(context.Context, *amqp.Message) error
	// }{
	// 	{"contact-created-topic", "contact-created-sub", consumers.HandleContactCreated},
	// 	{"contact-updated-topic", "contact-updated-sub", consumers.HandleContactUpdated},
	// 	{"upra-record-created-topic", "upra-record-created-sub", consumers.HandleUpraRecordCreated},
	// 	{"upra-record-deleted-topic", "upra-record-deleted-sub", consumers.HandleUpraRecordDeleted},
	// }

	// // Consume messages from each topic
	// for _, topicConfig := range topics {
	// 	go consumeFromTopic(client, topicConfig.topicName, topicConfig.subscriptionName, topicConfig.handler)
	// }

	// // Keep the main function running
	// select {}
}
