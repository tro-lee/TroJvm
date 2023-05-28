package heap

import (
	"TroJvm/classfile"
	_ "TroJvm/classfile"
)

// Class 定义类结构体
type Class struct {
	accessFlags       uint16        //访问标志
	name              string        //类名，完全限定名
	superClassName    string        //超类名
	interfaceNames    []string      //接口名
	constantPool      *ConstantPool //运行时常量池指针
	fields            []*Field      //字段表, 类成员
	methods           []*Method     //方法表，类成员
	loader            *ClassLoader  //读取类数据的类加载器
	superClass        *Class        //超类
	interfaces        []*Class      //接口
	instanceSlotCount uint          //实例变量占据的空间大小
	staticSlotCount   uint          //静态变量占据的空间大小
	staticVars        *Slots        //静态变量
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

// 用来判断某个访问标志符被设置，通过与运算判断是否等号

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
