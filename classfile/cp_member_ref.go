package classfile

// ConstantMemberrefInfo 字段或方法的符号引用
type ConstantMemberrefInfo struct {
	cp               ConstantPool // 常量池
	classIndex       uint16       // 类或接口的符号引用
	nameAndTypeIndex uint16       // 字段或方法的符号引用
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()       // 读取类或接口的符号引用
	self.nameAndTypeIndex = reader.readUint16() // 读取字段或方法的符号引用
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex) // 通过类或接口的符号引用索引获取类名
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex) // 通过字段或方法的符号引用索引获取字段或方法名和描述符
}

// ConstantFieldrefInfo 字段符号引用
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }

// ConstantMethodrefInfo 方法符号引用
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }

// ConstantInterfaceMethodrefInfo 接口方法符号引用
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }
