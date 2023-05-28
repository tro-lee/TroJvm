package heap

// SymRef 符号引用
type SymRef struct {
	cp        *ConstantPool // 常量池，运行时常量指针
	className string        // 类名，完全限定名
	class     *Class        // 类
}
