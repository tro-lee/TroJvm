package heap

import "TroJvm/classfile"

// FieldRef 定义字段引用
type FieldRef struct {
	MemberRef
	field *Field //指向属性
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
