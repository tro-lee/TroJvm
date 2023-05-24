package classfile

// ConstantNameAndTypeInfo 字段或方法
// 描述符分为类型描述符、字段描述符和方法描述符：
// 类型描述符：基本类型byte、short、char、int、long、float、double、boolean的描述符是单个字母，对象类型是L＋类的完全限定名＋分号，数组类型是[＋数组元素类型描述符
// 字段描述符：字段类型的描述符
// 方法描述符: (参数类型描述符)+返回值类型描述符
// 通过CONSTANT_NameAndType_info常量结构体同时包括名称和描述符，可以实现重载
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
