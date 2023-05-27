package control

import (
	base2 "TroJvm/instructions/base"
	"TroJvm/rtda"
)

// Branch always
type GOTO struct{ base2.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base2.Branch(frame, self.Offset)
}
