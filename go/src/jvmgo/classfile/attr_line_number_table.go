package classfile

// LineNumberTableAttribute LineNumberTable属性表, 存放方法的行号信息
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry // 行号表
}
type LineNumberTableEntry struct {
	startPc    uint16 // 起始PC
	lineNumber uint16 // 行号
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	// 读取行号表长度
	lineNumberTableLength := reader.readUint16()
	// 创建行号表
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	// 读取行号表
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
