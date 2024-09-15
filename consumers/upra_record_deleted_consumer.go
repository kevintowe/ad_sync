// consumers/upra_record_deleted_consumer.go
package consumers

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/go-amqp"
)

// HandleUpraRecordDeleted processes the UPRA Record Deleted event.
func HandleUpraRecordDeleted(ctx context.Context, msg *amqp.Message, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Processing UPRA Record Deleted: %s\n", string(msg.GetData()))

	// Simulate asynchronous processing
	// Example: Delete record from database.
}
