package rtda

// Thread 线程
// 每个线程都有一个pc寄存器和Java虚拟机栈
type Thread struct {
	pc    int    //pc寄存器
	stack *Stack //Java虚拟机栈的指针
}

// NewThread 生成线程
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// NewFrame 生成栈
func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return NewFrame(self, maxLocals, maxStack)
}

// Getter
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPc(pc int) {
	self.pc = pc
}

// JavaStack 的操作方法

// PushFrame 推入帧
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// PopFrame 推出帧
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// CurrentFrame 获取当前帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
