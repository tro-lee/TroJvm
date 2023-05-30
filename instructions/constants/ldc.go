package constants

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// LDC 把运行时常量池的常量加载，并把它推到操作数栈中
type LDC struct{ base.Index8Instruction }

// 大码
type LDC_W struct{ base.Index16Instruction }

// 大大码
type LDC2_W struct{ base.Index16Instruction }

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()

	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	default:
		panic("todo: ldc!")
	}
}

func (self *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

// Execute 对于long和double类型的常量，需要使用ldc2_w指令
func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()

	c := cp.GetConstant(uint(self.Index))
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
