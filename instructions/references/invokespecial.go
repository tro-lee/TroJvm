package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// INVOKE_SPECIAL 调用超类构造方法，实例初始化方法，私有方法
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
