package heap

// SymRef 符号引用
type SymRef struct {
	cp        *ConstantPool // 常量池，运行时常量指针
	className string        // 类名，完全限定名
	class     *Class        // 类
}

// ResolvedClass 类符号引用解析
func (self *SymRef) ResolvedClass() *Class {
	// 如果已经解析过了，就不解析了
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

// resolveClassRef 类符号引用解析
func (self *SymRef) resolveClassRef() {
	d := self.cp.class                      // 使用当前主类的类加载器
	c := d.loader.LoadClass(self.className) // 加载类
	// 类的访问控制规则
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
