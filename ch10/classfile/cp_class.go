package classfile

/*ConstantClassInfo {
	u1 tag;
	u2 name_index;
}
*/
type ConstantClassInfo struct {
	constantPool *ConstantPool
	nameIndex    uint16
}

func (mConstantClassInfo *ConstantClassInfo) readInfo(reader *ClassReader) {
	mConstantClassInfo.nameIndex = reader.readUint16()
}

// Name getter
func (mConstantClassInfo *ConstantClassInfo) Name() string {
	return mConstantClassInfo.constantPool.getUtf8(mConstantClassInfo.nameIndex)
}
