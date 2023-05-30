package heap

import (
	"TroJvm/classfile"
	_ "TroJvm/classfile"
	"strings"
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
	staticVars        Slots         //静态变量
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// newClass()方法把classFile转换成Class结构体
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

// NewObject 创建对象
func (self *Class) NewObject() *Object {
	return NewObject(self)
}

// GetMainMethod 调用类的静态方法
func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

// getStaticMethod 从方法表中查找方法
func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		// 判断方法是否是静态方法
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

// 访问控制
func (self *Class) isAccessibleTo(other *Class) bool {
	// 判断类是否可以访问
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

// 获取包名
func (self *Class) getPackageName() string {
	// 例如: github/good/Friend
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
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
