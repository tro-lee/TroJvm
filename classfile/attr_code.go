package classfile

// CodeAttribute 只存在method_info结构中，Code属性存放字节码等方法相关信息
type CodeAttribute struct {
	cp             ConstantPool           // 常量池
	maxStack       uint16                 // 操作数栈的最大深度
	maxLocals      uint16                 // 局部变量表大小
	code           []byte                 // 字节码
	exceptionTable []*ExceptionTableEntry // 异常处理表
	attributes     []AttributeInfo        // 属性表
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	// 读取操作数栈的最大深度
	self.maxStack = reader.readUint16()
	// 读取局部变量表大小
	self.maxLocals = reader.readUint16()
	// 读取字节码
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	// 读取异常处理表
	self.exceptionTable = readExceptionTable(reader)
	// 读取属性表
	self.attributes = readAttributes(reader, self.cp)
}

// ExceptionTableEntry 异常处理表
type ExceptionTableEntry struct {
	startPc   uint16 // try块的起始位置
	endPc     uint16 // try块的结束位置
	handlerPc uint16 // 异常处理代码的起始位置
	catchType uint16 // 捕获异常的类型
}

// readExceptionTable 读取异常处理表
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	// 异常处理表长度
	exceptionTableLength := reader.readUint16()
	// 创建异常处理表
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	// 遍历异常处理表
	for i := range exceptionTable {
		// 读取异常处理表
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	// 返回异常处理表
	return exceptionTable
}
