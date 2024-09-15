package main

import (
	"ad_sync/consumers"
	load_config "ad_sync/util"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Azure/go-amqp"
)

func main() {
	cfg := load_config.LoadConfig()

	// define the topics and their corresponding handlers
	topics := []struct {
		topicName        string
		subscriptionName string
		handler          func(context.Context, *amqp.Message) error
	}{
		{"contact-created-topic", "contact-created-sub", consumers.HandleContactCreated},
		{"contact-updated-topic", "contact-updated-sub", consumers.HandleContactUpdated},
		{"upra-record-created-topic", "upra-record-created-sub", consumers.HandleUpraRecordCreated},
		{"upra-record-deleted-topic", "upra-record-deleted-sub", consumers.HandleUpraRecordDeleted},
	}

	// Consume messages from each topic
	for _, topicConfig := range topics {
		go consumeFromTopic(client, topicConfig.topicName, topicConfig.subscriptionName, topicConfig.handler)
	}

	// Keep the main function running
	select {}
}
