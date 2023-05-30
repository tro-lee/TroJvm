package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
)

// PUT_STATIC 给类的某个静态变量赋值，使用两个操作数uint16，
// 找到字段符号引用，解析这个符号引用就能知道给哪个静态变量赋值，第二个操作数是赋值给静态变量的值，从操作栈中弹出
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	// 获得现在的方法区
	currentMethod := frame.Method()
	// 获取现在的类
	currentClass := currentMethod.Class()
	// 类的常量池
	cp := currentClass.ConstantPool()
	// 字段引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)

	// 获取字段
	field := fieldRef.ResolvedField()
	class := field.Class()

	// 发现不是静态变量抛出异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	//弹出操作栈的数进行操作，赋值给静态变量

	descriptor := field.Descriptor()
	slotId := field.SlotId()

	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
