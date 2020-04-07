package data

import "time"

type Message struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
	ReplyMe   bool      `json:"reply_me"`
	CreatedAt time.Time `json:"created_at"`
}
