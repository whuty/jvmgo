package classfile

// ConstantStringInfo sava an index to ConstantPool
// which point to a ConstantUtf8Info
type ConstantStringInfo struct {
	constantPool *ConstantPool
	stringIndex  uint16
}

func (mConstantStringInfo *ConstantStringInfo) readInfo(reader *ClassReader) {
	mConstantStringInfo.stringIndex = reader.readUint16()
}

func (mConstantStringInfo *ConstantStringInfo) String() string {
	return mConstantStringInfo.constantPool.getUtf8(mConstantStringInfo.stringIndex)
}
