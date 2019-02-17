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

func (q *Queue) Add(s Song, i int) error {
	if i == -1 {
		q.Songs = append(q.Songs, s)
	} else if i > len(q.Songs) {
		return errors.New("invalid i")
	} else {
		// insert song at a specific position
		q.Songs = append(q.Songs, Song{})
		copy(q.Songs[i+1:], q.Songs[i:])
		q.Songs[i] = s
	}

	return nil
}

func (q *Queue) Move(from, to int) error {
	if !q.isIndexValid(from) || !q.isIndexValid(to) {
		// TODO: specify which index is wrong
		return errors.New("invalid index")
	}

	// do nothing, indices are equal
	if from == to {
		return nil
	}

	s := q.Songs[from]
	q.Remove(from)
	q.Add(s, to)

	return nil
}

func (q *Queue) isIndexValid(index int) bool {
	return index >= 0 && index < len(q.Songs)
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
