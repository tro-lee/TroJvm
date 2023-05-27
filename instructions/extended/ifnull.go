package extended

import (
	base2 "TroJvm/instructions/base"
	"TroJvm/rtda"
)

// 判断是否为null，进行跳转
type IFNULL struct{ base2.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base2.Branch(frame, self.Offset)
	}
}

// 判断不是null，进行跳转
type IFNONNULL struct{ base2.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base2.Branch(frame, self.Offset)
	}
}
