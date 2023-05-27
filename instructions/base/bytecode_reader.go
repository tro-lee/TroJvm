package base

// BytecodeReader 字节码读取器
type BytecodeReader struct {
	// 存放code
	code []byte
	// 读取到哪个字节
	pc int
}

// Reset 重置，用来避免每次解码指令都新创建一个实例
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

// SkipPadding 跳过多余的字节，部分操作码的后面会有0~3字节的padding，保证defaultOffset在字节码中地址是4的倍数
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}

// ReadUint8 读取一个uint8整数
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

// ReadInt8 读取一个int8整数
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

// ReadUint16 读取一个uint16整数
func (self *BytecodeReader) ReadUint16() uint16 {
	// 先读取高8位
	byte1 := uint16(self.ReadUint8())
	// 再读取低8位
	byte2 := uint16(self.ReadUint8())
	// 拼接
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

// ReadInt32 读取一个int32整数
func (self *BytecodeReader) ReadInt32() int32 {
	// 先读取高16位
	byte1 := int32(self.ReadUint16())
	// 再读取低16位
	byte2 := int32(self.ReadUint16())
	// 拼接
	return (byte1 << 16) | byte2
}

// ReadInt32s 读取int32表
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	s := make([]int32, n)
	// 向后读取n个int32整数
	for i := range s {
		s[i] = self.ReadInt32()
	}
	return s
}
