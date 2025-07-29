package queue_test

import (
	"testing"

	"github.com/lewislk/go-utils/collection/queue"
)

func TestQueue(t *testing.T) {
	queue1 := queue.New[int]()
	queue1.Offer(1)
	t.Log(queue1.Poll())
	t.Log(queue1.Poll())
	t.Log("===============================分割线===============================")

	type Person struct {
		Name string `json:"name"`
	}
	p1 := Person{Name: "p1"}
	queue2 := queue.New[Person]()
	queue2.Offer(p1)
	t.Log(queue2.Poll())
	t.Log(queue2.Poll())
	t.Log("===============================分割线===============================")

	queue3 := queue.New[*Person]()
	queue3.Offer(&p1)
	t.Log(queue3.Poll())
	t.Log(queue3.Poll())
}
