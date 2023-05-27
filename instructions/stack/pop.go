package stack

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// 栈指令 直接对操作数栈进行操作

type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	// 弹出一个操作数栈
	frame.OperandStack().PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	// 弹出两个操作数栈
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}
