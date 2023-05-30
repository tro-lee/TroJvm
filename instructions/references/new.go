package references

import (
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
)

// NEW 创建类实例，操作uint16索引，
// 通过这个索引，从当前类的运行时常量池中找到一个类符号引用，解析类符号引用，拿到类数据，创建对象，把对象引用推入栈顶。
type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	// 获取栈帧所用方法的所属类，的常量池
	cp := frame.Method().Class().ConstantPool()
	// 获取这个引用
	ref := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 获取类
	class := ref.ResolvedClass()

	// 接口和抽象类不能实例化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	// 指向
	classRef := class.NewObject()
	// 操作数栈推入引用
	frame.OperandStack().PushRef(classRef)
}
