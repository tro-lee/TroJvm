package rtda

// Frame 栈帧
type Frame struct {
	lower        *Frame        //下一个栈帧
	localVars    LocalVars     //局部变量表，使用[]Slot实现
	operandStack *OperandStack //操作数栈，也使用[]Slot实现
	thread       *Thread       //当前所在线程
	nextPC       int           //下一个PC
}

func (f Frame) Thread() *Thread {
	return f.thread
}

// Getter
func (f Frame) NextPC() int {
	return f.nextPC
}

// Setter
func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func (f Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
