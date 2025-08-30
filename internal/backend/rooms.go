package backend

import "errors"

type Room struct {
	seats         map[int]*Agent
	messages      map[*Agent]*Message
	messages_pool chan *Message
}

func (r *Room) AppendAgent(a *Agent) (int, error) {
	for i := range r.seats {
		if r.seats[i] == nil {
			r.seats[i] = a
			return i, nil
		}
	}

	return -1, errors.New("table glitched")
}

func (r *Room) CheckAgent(a *Agent) error {
	if a.messages_queue
}