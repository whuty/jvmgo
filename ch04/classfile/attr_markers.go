package classfile

/*
DeprecatedAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
SyntheticAttribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

// MarkerAttribute has no data, so readInfo() method is empty
type MarkerAttribute struct{}

func (mAttribute *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
