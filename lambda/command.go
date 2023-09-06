package main

import (
	"log"
	"os/exec"
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
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) GenerateMessageId() *Command {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	c.MessageId = string(newUUID)
	return c
}

func (c *Command) WithActionByState(state string) *Command {
	if state == OPENED {
		c.Action = true
	} else {
		c.Action = false
	}

	return c
}
