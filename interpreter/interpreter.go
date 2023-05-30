package interpreter

import (
	"TroJvm/instructions"
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"TroJvm/rtda/heap"
	"fmt"
)

// Interpret 解释器
func Interpret(method *heap.Method) {
	// 创建一个线程
	thread := rtda.NewThread()
	// 根据方法创建一个帧
	frame := thread.NewFrame(method)

	thread.PushFrame(frame)
	defer catchErr(frame)

	loop(thread, method.Code())
}

// CatchErr 报错
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		//fmt.Println(frame.Thread().PC())
		//fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		//panic(r)
	}
}

// Loop 循环
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		//计算pc
		pc := frame.NextPC()
		thread.SetPc(pc)

		//解码指令
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//执行指令
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
