package heap

import (
	"TroJvm/classfile"
	"TroJvm/classpath"
	"fmt"
)

// ClassLoader 类加载器
type ClassLoader struct {
	cp       *classpath.ClassPath
	classMap map[string]*Class
}

// NewClassLoader 创建类加载器，给路径然后返回类加载器
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
	return self.loadNonArrayClass(name)
}

// 加载类
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 读取class文件
	data, entry := self.readClass(name)

	// 解析class
	class := self.defineClass(data)

	// 验证和准备
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)

	return class
}

// 读取class文件
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// 解析class
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 解析类
	class := parseClass(data)

	class.loader = self
	// 解析超类
	resolveSuperClass(class)
	// 解析接口
	resolveInterfaces(class)

	// 写入方法区
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	// Object没有超类
	if class.name != "java/lang/Object" {
		// 用它的类加载器，加载它的超类
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

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
	// 准备阶段，给类变量分配空间并给予初始值
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// prepare 准备阶段，给类变量分配空间并给予初始值
// 从上一个类的实例变量个数开始，给类的实例变量分配空间
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
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
	// 获取常量值索引
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
