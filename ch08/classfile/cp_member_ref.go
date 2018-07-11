package classfile

// fieldref info 表示字段符号引用
// methodref info 表示普通(非接口)方法符号引用
// interfaceref info 表示接口方法符号引用

/*
ConstantFieldrefInfo struct
*/
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }

/*
ConstantMethodrefInfo struct
*/
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }

/*
ConstantInterfaceMethodrefInfo struct
*/
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

/*
ConstantMemberrefInfo {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
*/
type ConstantMemberrefInfo struct {
	constantPool     *ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (mConstantMemberrefInfo *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	mConstantMemberrefInfo.classIndex = reader.readUint16()
	mConstantMemberrefInfo.nameAndTypeIndex = reader.readUint16()
}

//ClassName getter
func (mConstantMemberrefInfo *ConstantMemberrefInfo) ClassName() string {
	return mConstantMemberrefInfo.constantPool.getClassName(mConstantMemberrefInfo.classIndex)
}

//NameAndDescriptor getter
func (mConstantMemberrefInfo *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return mConstantMemberrefInfo.constantPool.getNameAndType(mConstantMemberrefInfo.nameAndTypeIndex)
}
