package extended

import (
	base2 "TroJvm/instructions/base"
	"TroJvm/rtda"
)

// GOTO_W 指令，goto指令，但是索引从2字节变成了4字节
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base2.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base2.Branch(frame, self.offset)
}
