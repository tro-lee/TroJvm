package extended

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// 判断是否为null，进行跳转
type IFNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

//判断不是null，进行跳转
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
