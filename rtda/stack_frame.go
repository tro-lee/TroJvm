package rtda

import "TroJvm/rtda/heap"

// Frame 栈帧
type Frame struct {
	lower        *Frame        //下一个栈帧
	localVars    LocalVars     //局部变量表，使用[]Slot实现
	operandStack *OperandStack //操作数栈，也使用[]Slot实现
	thread       *Thread       //当前所在线程
	nextPC       int           //下一个PC
	method       *heap.Method  //方法
}

func (f Frame) Method() *heap.Method {
	return f.method
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

func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}
