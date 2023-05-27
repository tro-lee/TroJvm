package base

import "TroJvm/go/src/jvmgo/rtda"

// Instruction 指令接口
type Instruction interface {
	// FetchOperands 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// Execute 执行指令逻辑
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction 无操作数指令
type NoOperandsInstruction struct {
}

func (i NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// noting to do
}

// BranchInstruction 跳转指令
type BranchInstruction struct {
	// 跳转偏移量
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	// 从字节码中读取一个uint16整数
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction 读取一个uint8整数
type Index8Instruction struct {
	// 局部变量表索引
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction 读取一个uint16整数
type Index16Instruction struct {
	// 局部变量索引
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
