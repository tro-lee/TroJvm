package loads

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// aload 操作引用类型变量，从局部变量表获取变量，然后推入操作数栈顶

type ALOAD struct{ base.Index8Instruction }

// 常用操作码

type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

// 从局部变量表获取变量，然后推入操作数栈顶
func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

// 相关执行方法
func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(self.Index))
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
