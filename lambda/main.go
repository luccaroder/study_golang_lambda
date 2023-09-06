package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.DynamoDBEvent) {
	iot := &IOT{}
	p := &Proxy{IotWrapper: iot}
	p.Proxy(ctx, event)
}

func main() {
	lambda.Start(handler)
}
