package classfile

// UnparsedAttribute 没有被解析的属性
type UnparsedAttribute struct {
	name   string // 属性名
	length uint32 // 属性长度
	info   []byte // 属性信息
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
