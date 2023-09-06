package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

type IOTWrapper interface {
	NewFromConfig(p client.ConfigProvider, cfgs ...*aws.Config) *iotdataplane.IoTDataPlane
}
