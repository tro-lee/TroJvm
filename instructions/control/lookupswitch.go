package control

import (
	base2 "TroJvm/instructions/base"
	"TroJvm/rtda"
)

// LOOKUP_SWITCH 指令，用于switch-case，用于不连续的key
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base2.BytecodeReader) {
	// 跳过padding()
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	// 弹出key，进行比较
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			base2.Branch(frame, int(offset))
			return
		}
	}
	base2.Branch(frame, int(self.defaultOffset))
}
