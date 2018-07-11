package classfile

/*
ExceptionsAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
变长属性记录方法抛出的异常表
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (mAttribute *ExceptionsAttribute) readInfo(reader *ClassReader) {
	mAttribute.exceptionIndexTable = reader.readUint16s()
}

// ExceptionIndexTable getter
func (mAttribute *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return mAttribute.exceptionIndexTable
}
