package constants

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// _ipush 从操作数中获取一个_ 型整数，扩展成int

// BIPUSH 获取byte类整数，扩展成int，并推入操作栈栈顶
type BIPUSH struct {
	val int8
}

// 继承指令接口

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// SIPUSH 获取short型整数，扩展成int，并推入操作数栈顶
type SIPUSH struct {
	val int16
}

// 继承指令接口

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
