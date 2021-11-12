// @Title		stack
// @Description	链栈实现。
// @Author		蔡与望
package stack

// 链栈结点。
type node struct {
	value interface{}
	next  *node
}

// 链栈结构。
type Stack struct {
	top   *node
	depth int
}

// 创建链栈。
func CreateStack() *Stack {
	return &Stack{
		top:   nil,
		depth: 0,
	}
}

// 获取栈顶元素值。
func (s *Stack) Top() interface{} {
	if s.depth == 0 {
		return nil
	}
	return s.top.value
}

// 获取栈深。
func (s *Stack) Depth() int {
	return s.depth
}

// 压栈。
func (s *Stack) Push(value interface{}) {
	node := &node{
		value: value,
		next:  s.top,
	}
	s.top = node
	s.depth++
}

// 弹栈。
func (s *Stack) Pop() interface{} {
	if s.depth == 0 {
		return nil
	}
	ret := s.top.value
	s.top = s.top.next
	s.depth--
	return ret
}
