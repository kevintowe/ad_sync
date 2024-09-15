package amqp

// consumeFromTopic connects to a service bus topic and subscription, and calls the appropriate handler for each message
func consumeFromTopic(host, sasKeyName, sasKey, topicName, subscriptionName string, handler func(context.Context, *amqp.Message, *sync.WaitGroup) ) {
	// create connection
	url := fmt.Sprintf("amqp://%s:%s@%s", sasKeyName, sasKey, host) // replace with ampq with ampqs
	client, err := amqp.Dial(url, amqp.ConnSASLPlain(sasKeyName, sasKey))
	if err != nil {
		log.Fatalf("failed to connect to AMQP broker: %v", err)
	}
	defer client.Close()

	// open a sesion
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create AMQP session: %v", err)
	}

	// construct the subscription receiver address
	subscriptionPath := fmt.Sprintf("%s/Subscriptions/%s", topicName, subscriptionName)

	// create receiver for the subscription
	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress(subscriptionPath)
		amqp.LinkCredit(10), // prefetch credit, what is this??
	)
	if err != nil {
		log.Fatalf("failed to create AMQP receiver: %v", err)
	}
	defer receiver.Close()

	// receive messages in a loop
	log.Printf("Listening for messages on topic: %s, subscription: %s", topicName, subscriptionName)
	for {
		ctx := context.Background()
		msg, err := receiver.Receive(ctx)
		if err != nil {
			log.Printf("failed to receive message: %v", err)
			break
		}

		// use a wait group to ensure the handler completes
		var wg sync.WaitGroup
		wg.Add(1)

		// call the handler asynchronously
		go func() {
			handler(ctx, msg, &wg)
			wg.Wait()

			// acknowledge the message after handler completes
			if err := msg.Accept(ctx); err != nil {
				log.Printf("failed to accept message: %v", err)
			}
		}()
	}
}
