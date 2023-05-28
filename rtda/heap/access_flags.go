package heap

const (
	ACC_PUBLIC       = 0x0001 //公有，用于类、字段、方法
	ACC_PRIVATE      = 0x0002 //私有，用于字段、方法
	ACC_PROTECTED    = 0x0004 //保护，用于字段、方法
	ACC_STATIC       = 0x0008 //静态，用于字段、方法
	ACC_FINAL        = 0x0010 //最终，用于类、字段、方法
	ACC_SUPER        = 0x0020 //超类，用于类
	ACC_SYNCHRONIZED = 0x0020 //同步，用于方法
	ACC_VOLATILE     = 0x0040 //易变，用于字段
	ACC_BRIDGE       = 0x0040 //桥接，用于方法
	ACC_TRANSIENT    = 0x0080 //瞬态，用于字段
	ACC_VARARGS      = 0x0080 //可变参数，用于方法
	ACC_NATIVE       = 0x0100 //本地，用于方法
	ACC_INTERFACE    = 0x0200 //接口，用于类
	ACC_ABSTRACT     = 0x0400 //抽象，用于类、方法
	ACC_STRICT       = 0x0800 //严格，用于方法
	ACC_SYNTHETIC    = 0x1000 //合成，用于类、字段、方法
	ACC_ANNOTATION   = 0x2000 //注解，用于类
	ACC_ENUM         = 0x4000 //枚举，用于类、字段
)
