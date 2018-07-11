package classfile

/*
constantPool_info{
	u1 tag;
	u1 info[];
}
*/

// the constant value of tag
const (
	CONSTANTClass              = 7
	CONSTANTFieldref           = 9
	CONSTANTMethodref          = 10
	CONSTANTInterfaceMethodref = 11
	CONSTANTString             = 8
	CONSTANTInteger            = 3
	CONSTANTFloat              = 4
	CONSTANTLong               = 5
	CONSTANTDouble             = 6
	CONSTANTNameAndType        = 12
	CONSTANTUtf8               = 1
	CONSTANTMethodHandle       = 15
	CONSTANTMethodType         = 16
	CONSTANTInvokeDynamic      = 18
)

// ConstantInfo interface
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANTInteger:
		return &ConstantIntegerInfo{}
	case CONSTANTFloat:
		return &ConstantFloatInfo{}
	case CONSTANTLong:
		return &ConstantLongInfo{}
	case CONSTANTDouble:
		return &ConstantDoubleInfo{}
	case CONSTANTUtf8:
		return &ConstantUtf8Info{}
	case CONSTANTString:
		return &ConstantStringInfo{constantPool: &cp}
	case CONSTANTClass:
		return &ConstantClassInfo{constantPool: &cp}
	case CONSTANTFieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{constantPool: &cp}}
	case CONSTANTMethodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{constantPool: &cp}}
	case CONSTANTInterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{constantPool: &cp}}
	case CONSTANTNameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANTMethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANTMethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANTInvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
