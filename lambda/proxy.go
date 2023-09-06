package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

type Proxy struct {
	IotWrapper IOTWrapper
}

func (controller *Proxy) Proxy(ctx context.Context, event events.DynamoDBEvent) {
	sess := session.Must(session.NewSession())

	iotDataClient := controller.IotWrapper.NewFromConfig(sess, aws.NewConfig().WithEndpoint("https://<IOT_API_ID>-ats.iot.us-east-1.amazonaws.com"))
	for _, record := range event.Records {
		fmt.Printf("Processing event ID %s, type %s\n", record.EventID, record.EventName)

		state := record.Change.NewImage["state"].String()

		iotMessage := NewCommand().GenerateMessageId().WithActionByState(state)

		iotByteMessage, err := json.Marshal(iotMessage)

		if err != nil {
			log.Printf("Error marshaling IoT message: %v", err)
			continue
		}

		publishInput := &iotdataplane.PublishInput{
			Topic:   aws.String("/iot/topic"),
			Qos:     aws.Int64(1),
			Payload: iotByteMessage,
		}

		_, err = iotDataClient.Publish(publishInput)
		if err != nil {
			log.Printf("Error publishing IoT message: %v", err)
		}
	}
}

func (iot *IOT) NewFromConfig(p client.ConfigProvider, cfgs ...*aws.Config) *iotdataplane.IoTDataPlane {
	return iotdataplane.New(p, cfgs...)
}
