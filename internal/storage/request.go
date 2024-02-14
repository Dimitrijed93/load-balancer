package storage

import "time"

type Request struct {
	Timestamp time.Time
	Service   string
}
