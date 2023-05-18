package model

import "time"

type Event struct {
	id          int        `sql:"id`
	title       string     `sql:"title"`
	description string     `sql:"description"`
	startedAt   *time.Time `sql:"started_at"`
	endedAt     *time.Time `sql:"ended_at"`
	releasedAt  *time.Time `sql:"released_at"`
	createdAt   *time.Time `sql:"created_at"`
}
