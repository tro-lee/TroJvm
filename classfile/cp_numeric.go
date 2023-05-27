package classfile

import "math"

// 字面量之一

// ConstantIntegerInfo 整数
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	// 读取数据
	bytes := reader.readUint32()
	// 转换数据
	self.val = int32(bytes)
}

// ConstantFloatInfo 浮点数
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	// 读取数据
	bytes := reader.readUint32()
	// 转换数据
	self.val = math.Float32frombits(bytes)
}

// ConstantLongInfo 长整数
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

// ConstantDoubleInfo 双精度浮点数
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
