package stack

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// dup 指令，用来复制栈顶的单个变量，不同的dup指令复制的方式不同

type DUP struct{ base.NoOperandsInstruction }
type DUP_X1 struct{ base.NoOperandsInstruction }
type DUP_X2 struct{ base.NoOperandsInstruction }
type DUP2 struct{ base.NoOperandsInstruction }
type DUP2_X1 struct{ base.NoOperandsInstruction }
type DUP2_X2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]

	\_
	  |
	  V

[...][c][b][a][a]
*/
func (self *DUP) Execute(frame *rtda.Frame) {
	// 赋值栈，便于操作
	stack := frame.OperandStack()
	// 弹出栈顶的变量
	slot := stack.PopSlot()
	// 复制变量
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
bottom -> top
[...][c][b][a]

	 __/
	|
	V

[...][c][a][b][a]
*/
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	// 赋值栈，便于操作
	stack := frame.OperandStack()
	// 弹出栈顶的变量
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	// 复制变量
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]

	 _____/
	|
	V

[...][a][c][b][a]
*/
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	// 赋值栈，便于操作
	stack := frame.OperandStack()
	// 弹出栈顶的变量
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	// 复制变量
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]____

	\____   |
	     |  |
	     V  V

[...][c][b][a][b][a]
*/
func (self *DUP2) Execute(frame *rtda.Frame) {
	// 赋值栈，便于操作
	stack := frame.OperandStack()
	// 弹出栈顶的变量
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	// 复制变量
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]

	 _/ __/
	|  |
	V  V

[...][b][a][c][b][a]
*/
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][d][c][b][a]

	 ____/ __/
	|   __/
	V  V

[...][b][a][d][c][b][a]
*/
func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
