// @Title: stack
// @Description: Implementation of stack with linked list.
// @Author: Yuwang Cai, Fumiama
package stack

import (
	"fmt"
	"strconv"
	"sync"
)

// Struct stack.
type Stack struct {
	data *[]interface{}
	p    uint
	mu   sync.RWMutex
}

// Create an empty stack.
func CreateStack() *Stack {
	return &Stack{
		data: new([]interface{}),
		p:    0,
	}
}

// Get top element value.
func (s *Stack) Top() interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.p == 0 {
		return nil
	}
	return (*s.data)[s.p-1]
}

// Get stack depth.
func (s *Stack) Depth() uint {
	s.mu.RLock()
	p := s.p
	s.mu.RUnlock()
	return p
}

// Push new element into stack.
func (s *Stack) Push(value interface{}) {
	s.mu.Lock()
	*s.data = append(*s.data, value)
	s.p++
	s.mu.Unlock()
}

// Pop out the top element.
func (s *Stack) Pop() interface{} {
	top := s.Top()
	if top != nil {
		s.mu.Lock()
		s.p--
		*s.data = (*s.data)[:s.p]
		s.mu.Unlock()
	}
	return top
}

// Pop out the top element as float64.
func (s *Stack) PopAsDouble() float64 {
	top := s.Pop()
	switch top := top.(type) {
	case float64:
		return top
	case string:
		topDouble, err := strconv.ParseFloat(top, 64)
		if err != nil {
			panic(err.Error())
		}
		return topDouble
	default:
		panic(fmt.Sprintf("popDouble() error: Invalid type %T", top))
	}
}
