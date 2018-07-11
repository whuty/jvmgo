package classfile

/*
ConstantNameAndTypeInfo {
	u1 tag;
	u2 name_index;
	u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (mConstantNameAndTypeInfo *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	mConstantNameAndTypeInfo.nameIndex = reader.readUint16()
	mConstantNameAndTypeInfo.descriptorIndex = reader.readUint16()
}
