package heap

import "TroJvm/classfile"

// InterfaceMethodRef 接口引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method //指向实现接口方法的方法
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
