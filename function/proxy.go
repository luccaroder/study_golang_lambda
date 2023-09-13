package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

type Proxy struct {
	IotWrapper IOTWrapper
}

func (controller *Proxy) Proxy(ctx context.Context, event events.DynamoDBEvent, config *AwsConfig) []error {
	iotDataClient := controller.IotWrapper.NewFromConfig(config.Session, config.Endpoint)

	var errs []error

	for _, record := range event.Records {
		fmt.Printf("Processing event ID %s, type %s\n", record.EventID, record.EventName)

		state := record.Change.NewImage["call"].Map()["currentStatus"].String()

		iotMessage := NewCommand("1234", "786855982", state)

		errs = iotMessage.Validate()

		if len(errs) > 0 {
			log.Printf("Error proccess command: %v", errs)
			continue
		}

		iotByteMessage, err := json.Marshal(iotMessage)
		if err != nil {
			log.Printf("Error marshaling IoT message: %v", err)
			continue
		}

		publishInput := &iotdataplane.PublishInput{
			Topic:   aws.String("/iot/dev/v1/alert/786855982/command"),
			Qos:     aws.Int64(1),
			Payload: iotByteMessage,
		}

		_, err = iotDataClient.Publish(publishInput)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (iot *IOT) NewFromConfig(p client.ConfigProvider, cfgs ...*aws.Config) *iotdataplane.IoTDataPlane {
	return iotdataplane.New(p, cfgs...)
}
