package math

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"math"
)

// rem 求余指令
type DREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type IREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	// 弹出栈顶两个int变量
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	// 出现除数为0的情况，抛出异常
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	// 计算求余，再推入
	result := v1 % v2
	stack.PushInt(result)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	// 弹出栈顶两个long变量
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	// 出现除数为0的情况，抛出异常
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	// 计算求余，再推入
	result := v1 % v2
	stack.PushLong(result)
}

func (self *FREM) Execute(frame *rtda.Frame) {
	// 弹出栈顶两个float变量
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	// 计算求余，再推入
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

func (self *DREM) Execute(frame *rtda.Frame) {
	// 弹出栈顶两个double变量
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	// 计算求余，再推入
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
