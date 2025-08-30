package backend

import (
	"errors"
	"math/rand"
	"time"
)

type Agent struct {
	id             int
	name           string
	messages_queue chan *Message
}

func NewAgent(name string) Agent {
	if name == "" {
		name = "New guest"
	}

	return Agent{rand.Intn(100000), name, make(chan *Message, 3)}
}

func (a *Agent) CreateMessage(s string) error {
	if s == "" {
		return errors.New("empty string input")
	}

	a.messages_queue <- &Message{a.id, time.Now(), s}
	return nil
}
