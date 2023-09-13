package main

import "errors"

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
		MessageId: MessageId,
		Action:    state == OPENED,
		AlertId:   alertId,
	}
}

func (c *Command) Validate() []error {
	var errs []error
	if c.MessageId == "" {
		errs = append(errs, errors.New("MessageId is required"))
	}

	if c.AlertId == "" {
		errs = append(errs, errors.New("AlertId is required"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// func (c *Command) GenerateMessageId() *Command {
// 	newUUID, err := exec.Command("uuidgen").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c.MessageId = string(newUUID)
// 	return c
// }

// func (c *Command) WithActionByState(state string) *Command {
// 	if state == OPENED {
// 		c.Action = true
// 	} else {
// 		c.Action = false
// 	}

// 	return c
// }

// func (c *Command) WithAlertId(alertId string) *Command {
// 	c.AlertId = alertId

// 	return c
// }
