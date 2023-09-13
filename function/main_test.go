package main

import (
	"context"
	"encoding/json"
	"testing"

	"lambda/golang/function/mocks"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/stretchr/testify/assert"
)

type errorTestCases struct {
	description string
	wantErr     []error
}

func TestHandler(t *testing.T) {
	input := []byte(`
	  { "M":
		  {
			  "currentStatus": { "S": "opened" }
		  }
	  }`)
	var av events.DynamoDBAttributeValue

	for _, tt := range []errorTestCases{
		{
			description: "Should return success when publish message",
			wantErr:     nil,
		},
	} {
		t.Run(tt.description, func(t *testing.T) {
			ctx := context.Background()

			err := json.Unmarshal(input, &av)

			assert.Nil(t, err)

			records := []events.DynamoDBEventRecord{
				{
					EventID:   "1",
					EventName: "INSERT",
					Change: events.DynamoDBStreamRecord{
						NewImage: map[string]events.DynamoDBAttributeValue{
							"call": av,
						},
					},
				},
				{
					EventID:   "1",
					EventName: "INSERT",
					Change: events.DynamoDBStreamRecord{
						NewImage: map[string]events.DynamoDBAttributeValue{
							"call": av,
						},
					},
				},
			}

			iot := mocks.NewIOTWrapper(t)
			sess := session.Must(session.NewSession())
			endpoint := aws.NewConfig().WithEndpoint("http://test-ats.iot.us-east-1.amazonaws.com")
			awsConfig := &AwsConfig{
				Session:  sess,
				Endpoint: endpoint,
			}

			iotPlane := &iotdataplane.IoTDataPlane{
				Client: &client.Client{Config: aws.Config{
					Endpoint: endpoint.Endpoint,
				}},
			}
			iot.On("NewFromConfig", sess, endpoint).Return(iotPlane)

			p := &Proxy{IotWrapper: iot}
			errs := p.Proxy(ctx, events.DynamoDBEvent{Records: records}, awsConfig)

			assert.Equal(t, tt.wantErr, errs)
		})
	}
}
