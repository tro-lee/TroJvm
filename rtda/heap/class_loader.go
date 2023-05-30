package heap

import (
	"TroJvm/classfile"
	"TroJvm/classpath"
	"fmt"
)

// ClassLoader 类加载器，关键让classpath和classfile和class联系起来
// 大致，基本把classpath的字节数据转换成classfile的数据，然后再转换成class的数据
type ClassLoader struct {
	cp       *classpath.ClassPath
	classMap map[string]*Class
}

// NewClassLoader 创建类加载器，给路径然后返回类加载器，梦的开始嗷
func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

// LoadClass 搜索和读取class文件
func (self *ClassLoader) LoadClass(name string) *Class {
	// 先判断是否已经加载过了
	if class, ok := self.classMap[name]; ok {
		// 类已经加载
		return class
	}
	// 正式加载类
	return self.loadNonArrayClass(name)
}

// 加载类，加载类三大步骤读取class文件，解析class，验证和准备(给内存)class
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 读取class文件，涉及到class文件的读取，所以需要classpath
	data, entry := self.readClass(name)

	// 解析class，使用classfile包解析class文件
	class := self.defineClass(data)

	// 验证和准备，重点给内存大小
	link(class)
	// 完成类加载
	fmt.Printf("[Loaded %s from %s]\n", name, entry)

	return class
}

// 读取class文件
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	// 使用classpath读取类数据，获取类的字节数据和entry（类路径项）
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// 解析class
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 解析类，调用classfile解析classpath传来的data数据，获得classfile结构体，然后转换成class
	class := parseClass(data)

	class.loader = self
	// 解析超类
	resolveSuperClass(class)
	// 解析接口
	resolveInterfaces(class)

	// 写入class_loader，作为缓存，下次类加载，发现已经加载了，就可以不加载了
	self.classMap[class.name] = class
	return class
}

// 解析类
func parseClass(data []byte) *Class {
	// 重点过程一：把class文件数据转换成class文件结构体
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	// 重点过程二：把class文件结构体转换成class结构体
	return newClass(cf)
}

// 解析超类
func resolveSuperClass(class *Class) {
	// Object没有超类
	if class.name != "java/lang/Object" {
		// 用它的类加载器，加载它的超类
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// 解析接口
func resolveInterfaces(class *Class) {
	// 获取接口数量
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 链接，进行验证和准备
func link(class *Class) {
	// 对类进行严格验证
	verify(class)
	// 准备阶段，计算出类的实例变量个数，给类的实例变量分配空间并给予初始值
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// prepare 准备阶段，给类变量分配空间并给予初始值
// 从上一个类的实例变量个数开始，给类的实例变量分配空间
func prepare(class *Class) {
	// 计算类的实例变量个数
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	// 给类变量分配空间并给予初始值，完成对class的静态变量表初始化
	allocAndInitStaticVars(class)
}

// 计算类的实例字段变量个数
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		// 等于超类的实例变量个数
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				field.slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// 计算类的静态变量个数
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 给类变量分配空间
func allocAndInitStaticVars(class *Class) {
	// 给类变量分配空间
	class.staticVars = newSlots(class.staticSlotCount)
	// 给类变量赋予初始值
	for _, field := range class.fields {
		// 给静态变量复制，从常量池中加载常量值
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 给静态变量赋予初始值
func initStaticFinalVar(class *Class, field *Field) {
	// 获取静态变量表
	vars := class.staticVars
	// 获取常量池
	cp := class.constantPool
	// 获取常量值索引，在类常量池转换时，顺便把常量值索引记录下来
	cpIndex := field.ConstValueIndex()
	// 获取静态变量索引
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.descriptor {
		case "Z", "B", "C", "S", "I":
			// 获取int类型的常量值
			val := cp.GetConstant(cpIndex).(int32)
			// 给静态变量赋值
			vars.SetInt(slotId, val)
		case "J":
			// 获取long类型的常量值
			val := cp.GetConstant(cpIndex).(int64)
			// 给静态变量赋值
			vars.SetLong(slotId, val)
		case "F":
			// 获取float类型的常量值
			val := cp.GetConstant(cpIndex).(float32)
			// 给静态变量赋值
			vars.SetFloat(slotId, val)
		case "D":
			// 获取double类型的常量值
			val := cp.GetConstant(cpIndex).(float64)
			// 给静态变量赋值
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
