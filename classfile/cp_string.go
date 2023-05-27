package classfile

// 字面量之一

// ConstantStringInfo 字符串
// 字符串本身不存放字符串数据，只存了常量池索引，指向CONSTANT_Utf8_info常量
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	// 读取索引
	self.stringIndex = reader.readUint16()
}
func (self *ConstantStringInfo) Name() string {
	return self.cp.getUtf8(self.stringIndex)
}
