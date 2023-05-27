package rtda

import "math"

// LocalVars 局部变量表
type LocalVars []Slot

// NewLocalVars 创建局部变量表
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// SetInt 设置int类型的局部变量
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

// GetInt 获取int类型的局部变量
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// SetFloat 设置float类型的局部变量
func (self LocalVars) SetFloat(index uint, val float32) {
	// 将float32类型的值转换成uint32类型的值
	bits := math.Float32bits(val)
	// 将uint32类型的值转换成int32类型的值
	self[index].num = int32(bits)
}

// GetFloat 获取float类型的局部变量
func (self LocalVars) GetFloat(index uint) float32 {
	// 将int32类型的值转换成uint32类型的值
	bits := uint32(self[index].num)
	// 将uint32类型的值转换成float32类型的值
	return math.Float32frombits(bits)
}

// SetLong 设置long类型的局部变量
func (self LocalVars) SetLong(index uint, val int64) {
	// 将int64类型的值转换成int32类型的值
	self[index].num = int32(val)
	// 将int64类型的值右移32位，然后转换成int32类型的值
	self[index+1].num = int32(val >> 32)
}

// GetLong 获取long类型的局部变量
func (self LocalVars) GetLong(index uint) int64 {
	// 将int32类型的值转换成int64类型的值
	low := uint32(self[index].num)
	// 将int32类型的值转换成int64类型的值
	high := uint32(self[index+1].num)
	// 将high左移32位，然后与low进行或运算
	return int64(high)<<32 | int64(low)
}

// SetDouble 设置double类型的局部变量
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

// GetDouble 获取double类型的局部变量
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

// SetRef 设置引用类型的局部变量
func (self LocalVars) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

// GetRef 获取引用类型的局部变量
func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}
