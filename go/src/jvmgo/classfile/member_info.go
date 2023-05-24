package classfile

// MemberInfo 字段和方法表
type MemberInfo struct {
	cp              ConstantPool    // 常量池
	accessFlags     uint16          //访问标志
	nameIndex       uint16          //常量池索引,给出字段名或方法名
	descriptorIndex uint16          //常量池索引,给出字段或方法描述符
	attributes      []AttributeInfo //属性表
}

// readMembers 生成读取字段表和方法表的函数
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	// 读取字段表或方法表的成员数量
	memberCount := reader.readUint16()
	// 生成字段表或方法表
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 辅助函数 readMember 用来读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	// 依次读取数据
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),        // 访问标志
		nameIndex:       reader.readUint16(),        // 名称索引
		descriptorIndex: reader.readUint16(),        // 描述符索引
		attributes:      readAttributes(reader, cp), // 属性表
	}
}

// Getter

// AccessFlags 获取访问标志
func (self *MemberInfo) AccessFlags() uint16 {}

// Name 获取名称
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// Descriptor 获取描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
