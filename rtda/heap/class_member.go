package heap

import "TroJvm/classfile"

// 类成员，包括字段和方法
type ClassMember struct {
	accessFlags uint16 //访问标志
	name        string //名字
	descriptor  string //标志
	class       *Class //所属类
}

// copyMemberInfo()方法从classfile.MemberInfo结构体中复制数据
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}
