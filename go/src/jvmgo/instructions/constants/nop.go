package constants

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// NOP Do nothing
type NOP struct{ base.NoOperandsInstruction }

// Execute 执行方法:什么都不做
func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么都不做
}
