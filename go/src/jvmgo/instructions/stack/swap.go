package stack

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// SWAP 交换栈顶两个元素
type SWAP struct{ base.NoOperandsInstruction }

// Execute 交换栈顶两个元素
func (self *SWAP) Execute(frame *rtda.Frame) {
	// 赋值栈，便于操作
	stack := frame.OperandStack()
	// 弹出栈顶的两个变量
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	// 交换变量
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
