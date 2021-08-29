package todo

import "time"

type Task struct {
	Id          int
	Description string
	Finished    bool
	FinishTime  time.Time
}
