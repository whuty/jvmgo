package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

/*
MemberInfo struct
*/
type MemberInfo struct {
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (mMemberInfo *MemberInfo) AccessFlags() uint16 {
	return mMemberInfo.accessFlags
}
func (mMemberInfo *MemberInfo) Name() string {
	return mMemberInfo.constantPool.getUtf8(mMemberInfo.nameIndex)
}
func (mMemberInfo *MemberInfo) Descriptor() string {
	return mMemberInfo.constantPool.getUtf8(mMemberInfo.descriptorIndex)
}

func (mMemberInfo *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mMemberInfo.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (mMemberInfo *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range mMemberInfo.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
