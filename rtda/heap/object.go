package heap

type Object struct {
	class *Class // 对象所属的类
	field Slots  // 对象的字段
}
