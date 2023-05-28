package heap

import "TroJvm/classfile"

// Field 字段
type Field struct {
	ClassMember
	slotId          uint // 字段的槽位索引
	constValueIndex uint // 静态变量的值在常量池中的索引
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	// 把数据复制到fields中
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	// 将classfile.MemberInfo中的constValueIndex属性复制到Field结构体中
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self *Field) SlotId() uint {
	return self.slotId
}
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}
