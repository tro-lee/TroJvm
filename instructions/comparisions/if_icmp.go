package comparisons

import (
	base2 "TroJvm/instructions/base"
	"TroJvm/rtda"
)

// IF_ICMP<cond> 指令，弹出两个int变量进行比较
type IF_ICMPEQ struct{ base2.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base2.Branch(frame, self.Offset)
	}
}

type IF_ICMPNE struct{ base2.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base2.Branch(frame, self.Offset)
	}
}

type IF_ICMPLT struct{ base2.BranchInstruction }

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base2.Branch(frame, self.Offset)
	}
}

type IF_ICMPLE struct{ base2.BranchInstruction }

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base2.Branch(frame, self.Offset)
	}
}

type IF_ICMPGT struct{ base2.BranchInstruction }

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base2.Branch(frame, self.Offset)
	}
}

type IF_ICMPGE struct{ base2.BranchInstruction }

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base2.Branch(frame, self.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
