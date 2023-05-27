package classfile

// AttributeInfo 属性信息接口
// 按照用户，23种预定义属性可分为三组，第一组是实现Java虚拟机所必需的，共有5种；第二组是Java类库所必需的，共有12种；第三组是可选的，共有6种。
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributes 读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	// 读取属性表长度
	attributesCount := reader.readUint16()
	// 创建属性表
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// readAttribute 读取属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 读取属性名索引
	attrNameIndex := reader.readUint16()
	// 在常量池中查询属性名
	attrName := cp.getUtf8(attrNameIndex)
	// 读取属性长度
	attrLen := reader.readUint32()
	// 创建属性信息
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	// 读取属性信息reader
	attrInfo.readInfo(reader)
	return attrInfo
}

// newAttributeInfo 创建属性信息 23种属性
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {

	// 标识的
	// Deprecated是最简单的属性之一，仅起到标记作用，不包含任何数据，可出现在ClassFile、field_info和method_info结构中。
	// Deprecated 用于表示不建议使用的属性，可以在java中使用@Deprecated标签指出编译器给出警告信息。
	case "Deprecated":
		return &DeprecatedAttribute{}
		// Synthetic也是最简单的属性之一，起到标识作用，不包含任何数据，可出现在ClassFile、field_info和method_info结构中。
		// Synthetic 用于标记源文件中不存在、由编译器生成的类成员，引入Synthetic属性主要是为了支持嵌套类和嵌套接口。
	case "Synthetic":
		return &SyntheticAttribute{}

		// 编译器生成的
		// SourceFile是可选定长属性，只会出现在ClassFile结构中，用于指出源文件名。
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}

		// 存放信息
		// ConstantValue是定长属性，只会出现在field_info结构中，用于表示常量表达式的值。
	case "ConstantValue":
		return &ConstantValueAttribute{}
		// Code是变长属性，只会出现在method_info结构中，用于存放字节码等方法相关信息。
	case "Code":
		return &CodeAttribute{cp: cp}
		// Exceptions是变长属性，用于指出方法抛出的异常表。
	case "Exceptions":
		return &ExceptionsAttribute{}

	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
