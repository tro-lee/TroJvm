package interpreter

import (
	"TroJvm/classfile"
	"TroJvm/instructions"
	"TroJvm/instructions/base"
	"TroJvm/rtda"
	"fmt"
)

// Interpret 解释器
func Interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	// 局部属性表最大索引
	maxLocals := codeAttr.MaxLocals()
	// 操作栈深度
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	//创建一个Thread实例，创建帧并推入虚拟机栈顶，然后执行方法
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

// CatchErr 报错
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
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
