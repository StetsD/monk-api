package models

import "time"

type Event struct {
	title       string
	dateStart   time.Time
	dateEnd     time.Time
	description string
	userId      int64
}
