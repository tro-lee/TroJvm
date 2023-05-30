package heap

// isAssignableFrom 判断类是否可以赋值给另一个类
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}
	if !t.IsInterface() {
		// t不是接口，则s是t的子类
		return s.isSubClassOf(t)
	} else {
		// t是接口，则s实现了t
		return s.isImplements(t)
	}
}

// IsSubClassOf 判断类是否是另一个类的子类
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// IsImplements 判断类是否实现了某个接口
func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// IsSubInterfaceOf 判断接口是否继承了某个接口
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
