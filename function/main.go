package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func handler(ctx context.Context, event events.DynamoDBEvent) {
	iot := &IOT{}
	p := &Proxy{IotWrapper: iot}
	awsConfig := &AwsConfig{
		Session:  session.Must(session.NewSession()),
		Endpoint: aws.NewConfig().WithEndpoint("https://<id_api>-ats.iot.us-east-1.amazonaws.com"),
	}
	p.Proxy(ctx, event, awsConfig)
}

func main() {
	lambda.Start(handler)
}
