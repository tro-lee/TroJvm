package classfile

// ConstantClassInfo 类或接口的符号引用
// 类和超类索引，以及接口表中接口索引索引都是指向CONSTANT_Class_info常量
type ConstantClassInfo struct {
	cp        ConstantPool // 常量池
	nameIndex uint16       // 常量池索引,指向CONSTANT_Utf8_info常量
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	// 读取索引
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
