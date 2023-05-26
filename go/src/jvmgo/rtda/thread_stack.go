package rtda

// Stack 栈
type Stack struct {
	// 最大容量
	maxSize uint
	// 当前大小
	size uint
	// 栈顶
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// push 推入
func (self *Stack) push(frame *Frame) {
	// 栈溢出
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	// 栈顶不为空，就存入栈顶
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

// pop 弹出
func (self *Stack) pop() *Frame {
	// 栈为空
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	// 弹出
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

// top 返回栈顶，但并不弹出
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
