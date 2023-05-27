package classfile

// ConstantUtf8Info MUTF-8编码的字符串
// MUTF-8和UTF-8的区别是null字符会被编码成2字节:0xC0、0x80，补充字符拆分为代理对分别编码
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	//获取长度
	length := uint32(reader.readUint16())
	//取出[]byte
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
