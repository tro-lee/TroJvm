package heap

import "TroJvm/classfile"

// MemberRef 字段符号引用和方法符号引用共有的信息
type MemberRef struct {
	SymRef
	name       string // 字段名
	descriptor string // 字段描述符，代表在java虚拟机的角度，一个类可以有多个同名字段
}

// copyMemberRefInfo()方法从classfile.MemberInfo结构体中复制数据
func (self *MemberRef) copyMemberRefInfo(memberInfo *classfile.ConstantMemberrefInfo) {
	self.className = memberInfo.ClassName()
	self.name, self.descriptor = memberInfo.NameAndDescriptor()
}
