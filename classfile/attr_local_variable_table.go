package classfile

// LocalVariableTableAttribute LocalVariableTable属性表，存放方法的局部变量信息
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry // 局部变量表
}

type LocalVariableTableEntry struct {
	startPc         uint16 // 起始PC
	length          uint16 // 长度
	nameIndex       uint16 // 名称索引
	descriptorIndex uint16 // 描述符索引
	index           uint16 // 索引
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	// 读取局部变量表长度
	localVariableTableLength := reader.readUint16()
	// 创建局部变量表
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	// 读取局部变量表
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
