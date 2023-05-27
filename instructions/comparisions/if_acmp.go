package comparisons

import (
	base2 "TroJvm/instructions/base"
	"TroJvm/rtda"
)

// IF_ACMPEQ和IF_ACMPNE指令把栈顶的两个引用弹出，根据引用是否相同进行跳转
type IF_ACMPEQ struct{ base2.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base2.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ base2.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base2.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
