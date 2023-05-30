package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
)

// PUT_FIELD 获取实例变量
type PUT_FIELD struct{ base.Index16Instruction }

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	// 获取当前方法、当前类、当前常量池
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)

	field := fieldRef.ResolvedField()

	// 如果是常量就抛出异常
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 如果是final修饰的，只能在构造函数中初始化
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			// OHHHHH!!!!空指针异常，打卡地
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	}
}
