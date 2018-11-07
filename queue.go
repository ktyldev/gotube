package main

import (
	"errors"
	"fmt"
	"log"
)

var _queue Queue = Queue{make([]Song, 0)}

type Queue struct {
	Songs []Song
}

func GetQueue() *Queue {
	return &_queue
}

func (q *Queue) Top() (Song, error) {
	if q.IsEmpty() {
		return Song{}, errors.New("queue is empty")
	}

	return q.Songs[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.Songs) == 0
}

func (q *Queue) Add(s Song) {
	q.Songs = append(q.Songs, s)
	log.Println(q.Songs)
}

func (q *Queue) Next() error {
	if q.IsEmpty() {
		return errors.New("queue is empty")
	}

	q.Songs = q.Songs[1:]
	log.Println("bump")

	return nil
}

func (q *Queue) Clear() {
	q.Songs = make([]Song, 0)
	log.Println("clear")
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
