package heap

import "TroJvm/classfile"

// Method 方法
type Method struct {
	ClassMember
	maxStack  uint   // 操作数栈的最大深度
	maxLocals uint   // 局部变量表大小
	code      []byte // 字节码
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	// 把数据复制到methods中
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	// 将classfile.MemberInfo中的maxStack、maxLocals、code属性复制到Method结构体中
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}
