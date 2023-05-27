package rtda

// Frame 栈帧
type Frame struct {
	lower        *Frame        //下一个栈帧
	localVars    LocalVars     //局部变量表，使用[]Slot实现
	operandStack *OperandStack //操作数栈，也使用[]Slot实现
}

// Getter
func (f Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
