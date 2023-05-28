package heap

import "TroJvm/classfile"

// 运行时常量池
// 主要存放两类信息：字面量和符号引用
// 字面量包括整数、浮点数和字符串字面量
// 符号引用包括类符号引用、字段符号引用、方法符号引用和接口方法符号引用

type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	// 将class文件中的常量池，转换为运行时常量池
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	// cp的引用类型，继承关系如下:
	// symref -> classRef
	//		  -> memberRef
	//						-> fieldRef
	//						-> methodRef

	for i := 0; i < cpCount; i++ {
		// 获取常量池信息
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			// 整数
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			// 浮点数
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			// 长整数，因为占两个位置，所以需要跳过一个位置
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			// 双精度浮点数，因为占两个位置，所以需要跳过一个位置
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			// 字符串字面量
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.Name()
		case *classfile.ConstantClassInfo:
			// 类符号引用
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			// 字段符号引用
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			// 方法符号引用
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			// 接口方法符号引用
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		}
	}

	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	// 获取常量池中指定索引的常量
	if c := self.consts[index]; c != nil {
		return c
	}
	panic("Invalid constant pool index!")
}
