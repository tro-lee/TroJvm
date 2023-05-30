package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
)

// GET_STATIC 获取类的静态变量
type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	// 获取当前常量池
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)

	// 解析获取字段
	field := fieldRef.ResolvedField()
	class := field.Class()

	// 如果不是静态变量抛出异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 获取类的静态变量
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	// 根据类型获取静态变量
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
