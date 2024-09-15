package consumers

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/go-amqp"
)

// HandleUpraRecordCreated processes the UPRA Record Created event.
func HandleUpraRecordCreated(ctx context.Context, msg *amqp.Message, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Processing UPRA Record Created: %s\n", string(msg.GetData()))

	// Simulate asynchronous processing
	// Example: Log event or notify another service.
}
