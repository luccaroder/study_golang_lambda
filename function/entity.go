package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type IOT struct{}

type AwsConfig struct {
	Session  *session.Session
	Endpoint *aws.Config
}
