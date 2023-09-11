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

func (controller *Proxy) Proxy(ctx context.Context, event events.DynamoDBEvent, config *AwsConfig) (string, error) {
	iotDataClient := controller.IotWrapper.NewFromConfig(config.Session, config.Endpoint)

	var errorsConcat error

	var sent string

	for _, record := range event.Records {
		fmt.Printf("Processing event ID %s, type %s\n", record.EventID, record.EventName)

		state := record.Change.NewImage["call"].Map()["currentStatus"].String()

		iotMessage := NewCommand().GenerateMessageId().WithActionByState(state).WithAlertId("786855982")

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
			if errorsConcat == nil {
				errorsConcat = err
			} else {
				errorsConcat = fmt.Errorf("%v, %v", errorsConcat, err)
			}
		} else {
			if sent == "" {
				sent = fmt.Sprintf("Message sent to %v", iotMessage.AlertId)
			} else {
				sent = fmt.Sprintf("%v, Message sent to %v", sent, iotMessage.AlertId)
			}
		}
	}

	return sent, errorsConcat
}

func (iot *IOT) NewFromConfig(p client.ConfigProvider, cfgs ...*aws.Config) *iotdataplane.IoTDataPlane {
	return iotdataplane.New(p, cfgs...)
}
