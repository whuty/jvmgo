package classfile

/*
ConstantMethodHandleInfo {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (mConstantMethodHandleInfo *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	mConstantMethodHandleInfo.referenceKind = reader.readUint8()
	mConstantMethodHandleInfo.referenceIndex = reader.readUint16()
}

/*
ConstantMethodTypeInfo {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (mConstantMethodTypeInfo *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	mConstantMethodTypeInfo.descriptorIndex = reader.readUint16()
}

/*
ConstantInvokeDynamicInfo {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (mConstantInvokeDynamicInfo *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	mConstantInvokeDynamicInfo.bootstrapMethodAttrIndex = reader.readUint16()
	mConstantInvokeDynamicInfo.nameAndTypeIndex = reader.readUint16()
}
