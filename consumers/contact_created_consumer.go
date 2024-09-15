package consumers

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/go-amqp"
)

// HandleContactCreated processes the Contact Created event.
func HandleContactCreated(ctx context.Context, msg *amqp.Message, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Processing Contact Created: %s\n", string(msg.GetData()))

	// Simulate asynchronous processing
	// Example: Save to database, send an email, etc.
	// Your actual processing logic goes here.
}
