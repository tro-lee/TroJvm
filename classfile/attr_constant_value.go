package classfile

// ConstantValueAttribute 常量值属性，只出现在field_info结构中，用于表示常量表达式的值。
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

// 获取常量值索引
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
