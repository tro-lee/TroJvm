package classfile

// ConstantPool 常量池，常量的表
type ConstantPool []ConstantInfo

// readConstantPool 读取常量池
// 需要注意的是表头给出的常量大小比实际大1
// 0是无效索引，不指向任何常量
// CONSTANT_Long_info和CONSTANT_Double_info各占两个位置，实际常量数量比n-1还小
func readConstantPool(reader *ClassReader) ConstantPool {
	// 读取常量池大小
	cpCount := int(reader.readUint16())
	// 生成常量池
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		// 根据类型进行操作
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

// getConstantInfo 按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	// 索引非法
	panic("Invalid constant pool index!")
}

// getNameAndType 在常量池查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// getClassName 在常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// getUtf8 读取utf8形式的字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
