package classfile

/*
UnparsedAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (mAttribute *UnparsedAttribute) readInfo(reader *ClassReader) {
	mAttribute.info = reader.readBytes(mAttribute.length)
}

// Info getter
func (mAttribute *UnparsedAttribute) Info() []byte {
	return mAttribute.info
}
