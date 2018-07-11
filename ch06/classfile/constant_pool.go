package classfile

// ConstantPool array of ConstantInfo
type ConstantPool []ConstantInfo

// the count read from constant pool n is 1 larger than the size of constant pool
// the constant pool index is from 1 to n-1, index 0 point to nothing
// ConstantLongInfo and ConstantDoubleInfo take 2 places
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (mConstantPool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := mConstantPool[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (mConstantPool ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := mConstantPool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := mConstantPool.getUtf8(ntInfo.nameIndex)
	_type := mConstantPool.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (mConstantPool ConstantPool) getClassName(index uint16) string {
	classInfo := mConstantPool.getConstantInfo(index).(*ConstantClassInfo)
	return mConstantPool.getUtf8(classInfo.nameIndex)
}

func (mConstantPool ConstantPool) getUtf8(index uint16) string {
	utf8Info := mConstantPool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
