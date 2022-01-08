package lib

import (
	storm "github.com/asdine/storm/v3"
)

type Task struct {
	ID    int    `storm:"id,increment"`
	Title string `storm:"index"`
}

func (t Task) String() string {
	return t.Title
}

func InitDB() (*storm.DB, error) {
	db, err := storm.Open("tasks.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
