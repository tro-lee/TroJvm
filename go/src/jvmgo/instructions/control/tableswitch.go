package control

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// TABLE_SWITCH 用于实现 switch-case 语句，通过index进行跳转，用于连续的case
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32 //索引表，存放high-low+1个int值
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// 跳过padding
	reader.SkipPadding()

	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()

	//获取需要跳转的偏移量
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	// 从操作数栈中弹出一个int变量，看他是否在范围之内
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	base.Branch(frame, offset)
}
