package data_structures

import (
	"errors"
	"fmt"
)

type Stack interface {
	Push(n int)
	Pop() (int, error)
	IsEmpty() bool
}

type ArrStack struct {
	storage []int
	head    int
}

func (s *ArrStack) Push(n int) {
	if s.head == len(s.storage) {
		s.storage = append(s.storage, n)
	} else {
		s.storage[s.head] = n
	}
	s.head += 1
}

func (s *ArrStack) Pop() (int, error) {
	if s.head == 0 {
		return 0, errors.New("cannot pop empty s")
	}
	if float64(len(s.storage))/float64(s.head) > 1.61 {
		var tmp = make([]int, int(float64(len(s.storage))/1.61)+1)
		copy(tmp, s.storage)
		s.storage = tmp
	}
	var res = s.storage[s.head-1]
	s.head -= 1
	return res, nil
}

func (s *ArrStack) IsEmpty() bool {
	if s.head == 0 {
		return true
	}
	return false
}

func (s *ArrStack) String() string {
	return fmt.Sprint(s.storage[:s.head])
}
