package rtda

// Slot 局部变量，用来实现局部变量表
type Slot struct {
	num int32   //存放整数，将基础变量类型转换成int32类型
	ref *Object //存放引用，将引用类型转换成*Object类型
}
