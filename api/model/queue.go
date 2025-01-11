package model

import "time"

// QueueEntry represents an entry in the processing queue stored in Redis.
type QueueEntry struct {
	UserId  int       `json:"user_id"`
	NrQueue int64     `json:"nr_queue"`
	Date    time.Time `json:"date"`
	Status  string    `json:"status"`
}
