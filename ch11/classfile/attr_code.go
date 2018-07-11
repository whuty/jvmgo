package classfile

/*
CodeAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	constantPool   ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (mAttribute *CodeAttribute) readInfo(reader *ClassReader) {
	mAttribute.maxStack = reader.readUint16()
	mAttribute.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	mAttribute.code = reader.readBytes(codeLength)
	mAttribute.exceptionTable = readExceptionTable(reader)
	mAttribute.attributes = readAttributes(reader, mAttribute.constantPool)
}

// MaxStack getter
// 操作数栈的最大深度
func (mAttribute *CodeAttribute) MaxStack() uint {
	return uint(mAttribute.maxStack)
}

// MaxLocals getter
// 局部变量表大小
func (mAttribute *CodeAttribute) MaxLocals() uint {
	return uint(mAttribute.maxLocals)
}

// Code getter
// 字节码
func (mAttribute *CodeAttribute) Code() []byte {
	return mAttribute.code
}

func (mAttribute *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range mAttribute.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

// ExceptionTable getter
// 异常处理表
func (mAttribute *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return mAttribute.exceptionTable
}

// ExceptionTableEntry struct
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

// StartPc of Exception getter
func (mAttribute *ExceptionTableEntry) StartPc() uint16 {
	return mAttribute.startPc
}

// EndPc of Exception getter
func (mAttribute *ExceptionTableEntry) EndPc() uint16 {
	return mAttribute.endPc
}

// HandlerPc of Exception getter
func (mAttribute *ExceptionTableEntry) HandlerPc() uint16 {
	return mAttribute.handlerPc
}

// CatchType of Exception getter
func (mAttribute *ExceptionTableEntry) CatchType() uint16 {
	return mAttribute.catchType
}
