package classfile

// MarkerAttribute 标记属性
type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// 不读取数
}

// Deprecated 表示弃用
type DeprecatedAttribute struct{ MarkerAttribute }

// Synthetic 表示由编译器生成的类成员
type SyntheticAttribute struct{ MarkerAttribute }
