package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
)

// INSTANCE_OF 判断是否为某类的实例，需要两个索引，然后在类的常量池找到类符号引用，解析符号引用，判断是否可以转换成该类
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	// 获取栈弹出的对象引用
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		// 如果是null，则把0推入操作数栈顶
		stack.PushInt(0)
		return
	}

	// 找到类符号引用，解析符号引用，判断是否可以转换成该类
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
