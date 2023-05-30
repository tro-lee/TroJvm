package heap

import "TroJvm/classfile"

// MemberRef 字段符号引用和方法符号引用共有的信息
type MemberRef struct {
	SymRef
	name       string // 字段名
	descriptor string // 字段描述符，代表在java虚拟机的角度，一个类可以有多个同名字段
}

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}

// copyMemberRefInfo()方法从classfile.MemberInfo结构体中复制数据
func (self *MemberRef) copyMemberRefInfo(memberInfo *classfile.ConstantMemberrefInfo) {
	self.className = memberInfo.ClassName()
	self.name, self.descriptor = memberInfo.NameAndDescriptor()
}

// 访问控制规则
func (self *Field) isAccessibleTo(d *Class) bool {
	// 如果是public，就是任意访
	if self.IsPublic() {
		return true
	}
	c := self.class

	// 如果是protected, 就是同一个类，子类，同一个包
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}

	// 如果是private, 就是同一个类
	if self.IsPrivate() {
		return d == c
	}

	return c.getPackageName() == d.getPackageName()
}
