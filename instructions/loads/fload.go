package loads

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// fload 操作float变量，从局部变量表获取变量，然后推入操作数栈顶

type FLOAD struct{ base.Index8Instruction }

// 常用操作码

type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

// 从局部变量表获取变量，然后推入操作数栈顶
func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

// 相关执行方法
func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, uint(self.Index))
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
