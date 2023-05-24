package classfile

import "fmt"

// ClassFile  class文件结构体
type ClassFile struct {
	minorVersion uint16 // 次版本号
	majorVersion uint16 // 主版本号

	constantPool ConstantPool // 常量池

	accessFlags uint16 // 访问标志
	thisClass   uint16 // 类索引
	superClass  uint16 // 父类索引

	interfaces []uint16 // 接口索引表

	fields  []*MemberInfo // 字段表
	methods []*MemberInfo // 方法表

	attributes []AttributeInfo // 属性表
}

// Parse 解析class文件
func Parse(classData []byte) (cf *ClassFile, err error) {
	// defer + recover做到异常捕获
	defer func() {
		// 尝试恢复，如果恢复成功，将err赋值为恢复的错误
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error) // 类型断言
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	// 读取数据
	cr := &ClassReader{classData}
	cf = &ClassFile{}

	// 装填数据
	cf.read(cr)
	return
}

// read 读取
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckVersion(reader)             // 读取并检查版本号
	self.readAndCheckMagic(reader)               // 读取并检查魔数
	self.constantPool = readConstantPool(reader) // 读取常量池

	self.accessFlags = reader.readUint16() // 读取访问标志
	self.thisClass = reader.readUint16()   // 读取类索引
	self.superClass = reader.readUint16()  // 读取父类索引
	self.interfaces = reader.readUint16s() // 读取接口索引表

	self.fields = readMembers(reader, self.constantPool)        // 读取字段表
	self.methods = readMembers(reader, self.constantPool)       // 读取方法表
	self.attributes = readAttributes(reader, self.constantPool) // 读取属性表
}

// 开始使用的两个方法，进行魔数与版本的检验

// readAndCheckMagic 读取并检查魔数，魔数就是固定字节的开头，起到标识作用
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	//读取4个字节
	magic := reader.readUint32()
	//class的魔数为0xCAFEBABE
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// readAndCheckVersion 读取并检查版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 此版本号与主版本号分别占2个字节
	// 每次有大的Java版本发布，主版本号加一,Oracle的虚拟机完全向后兼容
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	// 鼠鼠的是61
	case 61, 60, 59, 58, 57, 56, 55:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// 相当于Getter

// MinorVersion 次版本号
func (self *ClassFile) MinorVersion() uint16 { return self.majorVersion }

// MajorVersion 主版本号
func (self *ClassFile) MajorVersion() uint16 { return self.majorVersion }

// ConstantPool 常量池
func (self *ClassFile) ConstantPool() ConstantPool { return self.constantPool }

// AccessFlags 访问标志
func (self *ClassFile) AccessFlags() uint16 { return self.accessFlags }

// Fields 字段表
func (self *ClassFile) Fields() []*MemberInfo { return self.fields }

// Methods 方法表
func (self *ClassFile) Methods() []*MemberInfo { return self.methods }

// 运用到常量池

// ClassName 类名 从常量池里查询
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName 父类名 从常量池里查询
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

// InterfaceNames 接口名
func (self *ClassFile) InterfaceNames() []string {
	// 接口名表
	interfaceNames := make([]string, len(self.interfaces))

	// 遍历接口索引表
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}

	return interfaceNames
}
