package heap

type Object struct {
	class  *Class // 对象所属的类
	fields Slots  // 对象的字段
}

func (o Object) Class() *Class {
	return o.class
}

func (o Object) Fields() Slots {
	return o.fields
}

// NewObject 实例化类
func NewObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// IsInstanceOf 判断对象是否是某个类的实例
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
