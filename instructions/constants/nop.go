package constants

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
)

// NOP Do nothing
type NOP struct{ base.NoOperandsInstruction }

// Execute 执行方法:什么都不做
func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么都不做
}
