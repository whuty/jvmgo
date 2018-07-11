package classfile

/*
SourceFileAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
attribute length must be 2
sourcefile 是可选定长属性,只出现在ClassFile结构中,用于指出源文件名
如果不生成此属性, 抛出异常时不会显示出错代码所属文件名
*/
type SourceFileAttribute struct {
	constantPool    ConstantPool
	sourceFileIndex uint16
}

func (mAttribute *SourceFileAttribute) readInfo(reader *ClassReader) {
	mAttribute.sourceFileIndex = reader.readUint16()
}

// FileName of sourcefile
func (mAttribute *SourceFileAttribute) FileName() string {
	return mAttribute.constantPool.getUtf8(mAttribute.sourceFileIndex)
}
