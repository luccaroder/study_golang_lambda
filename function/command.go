package main

import (
	"errors"

	"github.com/google/uuid"
)

type ACTION int

const (
	_      ACTION = iota
	OPENED        = "opened"
	CLOSED        = "closed"
)

type Command struct {
	MessageId string `json:"messageId"`
	Action    bool   `json:"action"`
	AlertId   string `json:"alertId"`
}

func NewCommand(MessageId string, alertId string, state string) *Command {
	return &Command{
		MessageId: uuid.New().String(),
		Action:    state == OPENED,
		AlertId:   alertId,
	}
}

func (c *Command) IsValidAlertId() error {
	if c.AlertId == "" {
		return errors.New("AlertId is required")
	}
	return nil
}
