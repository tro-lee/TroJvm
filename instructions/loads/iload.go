package loads

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// iload 操作int变量，从局部变量表获取变量，然后推入操作数栈顶

type ILOAD struct{ base.Index8Instruction }

// 常用操作码

type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

// 从局部变量表获取变量，然后推入操作数栈顶
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

// 相关执行方法

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
