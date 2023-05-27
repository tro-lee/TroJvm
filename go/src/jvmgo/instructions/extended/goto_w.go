package extended

import (
	"TroJvm/go/src/jvmgo/instructions/base"
	"TroJvm/go/src/jvmgo/rtda"
)

// GOTO_W 指令，goto指令，但是索引从2字节变成了4字节
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
