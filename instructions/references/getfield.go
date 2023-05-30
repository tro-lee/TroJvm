package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
)

// GET_FIELD 获取对象的实例变量值，然后推入操作数栈顶
type GET_FIELD struct{ base.Index16Instruction }

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)

	field := fieldRef.ResolvedField()

	// 如果是静态字段，则抛出IncompatibleClassChangeError异常
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 弹出对象引用，如果是null，则抛出NullPointerException异常
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

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
