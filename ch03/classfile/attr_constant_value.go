package classfile

/*
ConstantValueAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
定长属性 length = 2
只会出现在field info 结构中, 用于表示常量表达式的值
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (mAttribute *ConstantValueAttribute) readInfo(reader *ClassReader) {
	mAttribute.constantValueIndex = reader.readUint16()
}

// ConstantValueIndex getter
func (mAttribute *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return mAttribute.constantValueIndex
}
