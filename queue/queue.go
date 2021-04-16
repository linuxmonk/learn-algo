package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type Queue struct {
	mux     sync.Mutex
	queue   []int
	delpos  int
	pos     int
	size    int
	maxSize int
}

func New(max int) *Queue {
	return &Queue{
		queue:   make([]int, max),
		maxSize: max,
	}
}

func (q *Queue) Enqueue(elem int) error {
	if q == nil {
		return errors.New("invalid queue")
	}
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.size == q.maxSize {
		return errors.New("queue is full")
	}
	q.queue[q.pos%q.maxSize] = elem
	q.pos = (q.pos + 1) % q.maxSize
	q.size++
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q == nil {
		return -1, errors.New("invalid queue")
	}
	q.mux.Lock()
	defer q.mux.Unlock()
	if q.size == 0 {
		return -1, errors.New("queue is empty")
	}
	elem := q.queue[q.delpos]
	q.delpos = (q.delpos + 1) % q.maxSize
	q.size--
	return elem, nil
}

func (q *Queue) Print() {
	if q == nil {
		fmt.Println("Invalid Queue")
		return
	}
	if q.size == 0 {
		fmt.Println("<empty>")
		return
	}
	fmt.Println("::: Queue :::")
	i := q.delpos
	n := q.size
	for n > 0 {
		fmt.Printf("%d ", q.queue[i])
		i = (i + 1) % q.maxSize
		n--
	}
	fmt.Println()
	fmt.Println("::: Queue :::")
}

func main() {
	q := New(5)
	for i := 0; i < 6; i++ {
		if err := q.Enqueue(i); err != nil {
			fmt.Println("Enqueue: ", err)
			break
		}
	}
	q.Print()
	elem, err := q.Dequeue()
	if err != nil {
		fmt.Println("Dequeue: ", err)
		os.Exit(1)
	}
	fmt.Println("Dequeue: ", elem)

	for i := 0; i < 2; i++ {
		elem, err = q.Dequeue()
		if err != nil {
			fmt.Println("Dequeue: ", err)
			continue
		}
		fmt.Println("Dequeue: ", elem)
	}
	q.Print()

	fmt.Println("Enqueue: 100")
	if err = q.Enqueue(100); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	q.Print()

	fmt.Println("Full Dequeue")
	for i := 0; i < 5; i++ {
		elem, err = q.Dequeue()
		if err != nil {
			fmt.Println("Dequeue: ", err)
			continue
		}
		fmt.Println("Dequeue: ", elem)
	}

	q.Print()

	fmt.Println("Insert again")
	for i := 0; i < 3; i++ {
		if err := q.Enqueue(i); err != nil {
			fmt.Println("Enqueue: ", err)
			break
		}
	}
	q.Print()
}
