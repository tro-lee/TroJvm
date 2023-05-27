package classfile

import "encoding/binary"

// ClassReader  作用是向后读取某长度的字节
type ClassReader struct {
	data []byte
}

// readUint8 读取u1类型数据 8比特无符号整数
func (self *ClassReader) readUint8() uint8 {
	// 获得数据
	val := self.data[0]
	// 截取后面数据
	self.data = self.data[1:]
	// 返回数据
	return val
}

// readUint16 读取u2类型数据 16比特无符号整数
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// readUint32 读取u4类型数据 32比特无符号整数
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// readUint64 读取u8类型数据 64比特无符号整数
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// readUint16s 读取uint16表
func (self *ClassReader) readUint16s() []uint16 {
	// 读取长度
	n := self.readUint16()
	// 创建数组
	s := make([]uint16, n)
	// 循环读取
	for i := range s {
		s[i] = self.readUint16()
	}
	// 返回数组
	return s
}

// readBytes 读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
