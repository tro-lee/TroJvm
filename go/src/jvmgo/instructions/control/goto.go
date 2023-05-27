package control

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
