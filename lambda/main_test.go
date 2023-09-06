package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("TestHandler", func(t *testing.T) {
		ctx := context.Background()
		records := []events.DynamoDBEventRecord{
			{
				EventID:   "1",
				EventName: "INSERT",
			},
		}

		handler(ctx, events.DynamoDBEvent{Records: records})

	})
}
