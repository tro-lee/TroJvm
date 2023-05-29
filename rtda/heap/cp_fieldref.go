package heap

import "TroJvm/classfile"

// FieldRef 定义字段引用
type FieldRef struct {
	MemberRef
	field *Field //指向属性
}

// ResolvedField 字段符号引用解析
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolvedFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolvedFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	// 查找字段
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	// 判断字段是否可以访问
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
}

// lookupField 查找字段
func lookupField(class *Class, name, descriptor string) *Field {
	// 先在当前类中查找
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	// 在接口中查找
	for _, iface := range class.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	// 在父类中查找
	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}
	return nil
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
