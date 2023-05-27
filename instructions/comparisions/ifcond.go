package comparisons

import (
	base2 "TroJvm/instructions/base"
	"TroJvm/rtda"
)

// IF的相关判断指令
type IFEQ struct{ base2.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base2.Branch(frame, self.Offset)
	}
}

type IFNE struct{ base2.BranchInstruction }

func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base2.Branch(frame, self.Offset)
	}
}

type IFLT struct{ base2.BranchInstruction }

func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base2.Branch(frame, self.Offset)
	}
}

type IFLE struct{ base2.BranchInstruction }

func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base2.Branch(frame, self.Offset)
	}
}

type IFGT struct{ base2.BranchInstruction }

func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base2.Branch(frame, self.Offset)
	}
}

type IFGE struct{ base2.BranchInstruction }

func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base2.Branch(frame, self.Offset)
	}
}
