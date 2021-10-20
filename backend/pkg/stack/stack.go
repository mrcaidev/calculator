// @Title: stack
// @Description: Implementation of stack with linked list.
// @Author: Yuwang Cai
package stack

import (
	"fmt"
	"strconv"
)

// Node in linked list.
type node struct {
	value interface{}
	next  *node
}

// Struct stack.
type Stack struct {
	top   *node
	depth int
}

// Create an empty stack.
func CreateStack() *Stack {
	return &Stack{
		top:   nil,
		depth: 0,
	}
}

// Get top element value.
func (s *Stack) Top() interface{} {
	if s.depth == 0 {
		return nil
	}
	return s.top.value
}

// Get stack depth.
func (s *Stack) Depth() int {
	return s.depth
}

// Push new element into stack.
func (s *Stack) Push(value interface{}) {
	node := &node{
		value: value,
		next:  s.top,
	}
	s.top = node
	s.depth++
}

// Pop out the top element.
func (s *Stack) Pop() interface{} {
	if s.depth == 0 {
		return nil
	}
	ret := s.top.value
	s.top = s.top.next
	s.depth--
	return ret
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
