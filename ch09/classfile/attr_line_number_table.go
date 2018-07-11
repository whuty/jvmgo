package classfile

/*
LineNumberTableAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
存放方法的行号信息
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

// LineNumberTableEntry struct, startPc, lineNumber
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (mAttribute *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	mAttribute.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range mAttribute.lineNumberTable {
		mAttribute.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

// GetLineNumber of method
func (mAttribute *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(mAttribute.lineNumberTable) - 1; i >= 0; i-- {
		entry := mAttribute.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
