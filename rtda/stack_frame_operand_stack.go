package rtda

import (
	"TroJvm/rtda/heap"
	"math"
)

// OperandStack 操作数栈
type OperandStack struct {
	size  uint   //栈顶位置
	slots []Slot //栈
}

// newOperandStack 创建操作数栈
func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// PushSlot 推入Slot
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

// PopSlot 推出Slot
func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

// PushInt 推入int
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

// PopInt 弹出int
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

// PushFloat 推入float
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

// PopFloat 弹出float
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

// PushLong 推入long
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

// PopLong 弹出long
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

// PushDouble 推入double
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

// PopDouble 弹出double
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

// PushRef 推入引用
func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].ref = ref
	self.size++
}

// PopRef 弹出引用
func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}
