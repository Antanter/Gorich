package backend

import "time"

type Message struct {
	agent_id int
	time     time.Time
	text     string
}
