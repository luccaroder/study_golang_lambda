package main

import (
	"context"
	"testing"

	"lambda/golang/function/mocks"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

		iot := mocks.NewIOTWrapper(t)
		sess := session.Must(session.NewSession())
		endpoint := aws.NewConfig().WithEndpoint("http://test-ats.iot.us-east-1.amazonaws.com")
		awsConfig := &AwsConfig{
			Session:  sess,
			Endpoint: endpoint,
		}
		iot.On("NewFromConfig", sess, endpoint).Return(nil)
		p := &Proxy{IotWrapper: iot}
		p.Proxy(ctx, events.DynamoDBEvent{Records: records}, awsConfig)
	})
}
