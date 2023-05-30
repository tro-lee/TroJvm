package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
)

// CHECK_CAST 也进行判断，但不改变操作数栈，只抛出异常

type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 弹出对象引用
	ref := stack.PopRef()
	// 如果是null，则直接推入null
	if ref == nil {
		return
	}

	// 否则，获取类符号引用，解析符号引用，判断是否可以转换成该类
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	// 如果不可以转换，则抛出ClassCastException异常
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
