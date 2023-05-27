package constants

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// const 系列指令，把常量值（操作码指出的数）推入操作数栈顶
// 在这里，我们定义了15条指令：
// ACONST_NULL, DCONST_0, DCONST_1, FCONST_0, FCONST_1,
// FCONST_2, ICONST_M1, ICONST_0, ICONST_1, ICONST_2,
// ICONST_3, ICONST_4, ICONST_5, LCONST_0, LCONST_1

// ACONST_NULL 把null引用推入操作数栈顶
type ACONST_NULL struct{ base.NoOperandsInstruction }

// Execute 执行方法:把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// DCONST_0 把double型0推入操作数栈顶
type DCONST_0 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把double型0推入操作数栈顶
func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// DCONST_1 把double型1推入操作数栈顶
type DCONST_1 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把double型1推入操作数栈顶
func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// FCONST_0 把float型0推入操作数栈顶
type FCONST_0 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把float型0推入操作数栈顶
func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

// FCONST_1 把float型1推入操作数栈顶
type FCONST_1 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把float型1推入操作数栈顶
func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

// FCONST_2 把float型2推入操作数栈顶
type FCONST_2 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把float型2推入操作数栈顶
func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// ICONST_M1 把int型-1推入操作数栈顶
type ICONST_M1 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把int型-1推入操作数栈顶
func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

// ICONST_0 把int型0推入操作数栈顶
type ICONST_0 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把int型0推入操作数栈顶
func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

// ICONST_1 把int型1推入操作数栈顶
type ICONST_1 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把int型1推入操作数栈顶
func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

// ICONST_2 把int型2推入操作数栈顶
type ICONST_2 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把int型2推入操作数栈顶
func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

// ICONST_3 把int型3推入操作数栈顶
type ICONST_3 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把int型3推入操作数栈顶
func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

// ICONST_4 把int型4推入操作数栈顶
type ICONST_4 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把int型4推入操作数栈顶
func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

// ICONST_5 把int型5推入操作数栈顶
type ICONST_5 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把int型5推入操作数栈顶
func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// LCONST_0 把long型0推入操作数栈顶
type LCONST_0 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把long型0推入操作数栈顶
func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

// LCONST_1 把long型1推入操作数栈顶
type LCONST_1 struct{ base.NoOperandsInstruction }

// Execute 执行方法:把long型1推入操作数栈顶
func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
