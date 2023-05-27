package classfile

// tag 常量值定义,JVM定义的
const (
	// 类
	CONSTANT_Class = 7
	// 字段引用
	COMSTANT_Fieldref = 9
	// 方法引用
	CONSTANT_Methodref = 10
	// 接口方法引用
	CONSTANT_InterfaceMethodref = 11
	// 字符串
	CONSTANT_String = 8
	// 整数
	CONSTANT_Integer = 3
	// 浮点数
	CONSTANT_Float = 4
	// 长整数
	CONSTANT_Long = 5
	// 双精度浮点数
	CONSTANT_Double = 6
	// 名字和类型
	CONSTANT_NameAndType = 12
	// utf8
	CONSTANT_Utf8 = 1
	// 方法句柄
	CONSTANT_MethodHandle = 15
	// 方法类型
	CONSTANT_MethodType = 16
	// 动态调用
	CONSTANT_InvokeDynamic = 17
)

// ConstantInfo 定义常量信息接口
type ConstantInfo interface {
	// 读取常量信息
	readInfo(reader *ClassReader)
}

// readConstantInfo 读取常量信息
// 先读出tag值，创建具体的常量，最后调用常量的readInfo()方法读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// 读出什么类型的
	tag := reader.readUint8()
	// 创建常量
	c := newConstantInfo(tag, cp)
	// 读取常量信息
	c.readInfo(reader)
	return c
}

// newConstantInfo 生成常量信息
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case COMSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
		// 目的解析invokedynamic指令
	//case CONSTANT_MethodType:
	//	return &ConstantMethodTypeInfo{}
	//case CONSTANT_MethodHandle:
	//	return &ConstantMethodHandleInfo{}
	//case CONSTANT_InvokeDynamic:
	//	return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
