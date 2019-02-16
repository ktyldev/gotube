package main

import (
	"errors"
	"fmt"
)

var _queue Queue = Queue{make([]Song, 0)}

type Queue struct {
	Songs []Song
}

func GetQueue() *Queue {
	return &_queue
}

func (q *Queue) IsEmpty() bool {
	return len(q.Songs) == 0
}

func (q *Queue) Add(s Song) {
	q.Songs = append(q.Songs, s)
}

func (q *Queue) Remove(index int) error {
	if index >= len(q.Songs) {
		return errors.New("invalid index")
	}

	// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-array-in-golang
	q.Songs = append(q.Songs[:index], q.Songs[index+1:]...)

	return nil
}

func (q *Queue) Clear() {
	if q.IsEmpty() {
		return
	}

	q.Songs = make([]Song, 0)
}

func (q *Queue) GetById(id string) (Song, error) {
	var result Song

	for _, s := range q.Songs {
		if s.Id == id {
			result = s
			break
		}
	}

	if result.Id == "" {
		msg := fmt.Sprintf("no song matching %s\n", id)
		return Song{}, errors.New(msg)
	}

	return result, nil
}
