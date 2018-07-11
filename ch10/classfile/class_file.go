package classfile

import (
	"fmt"
)

/*
ClassFile {
	u4				magic;
	u2				minor_version
	u2				marjor_version
	u2				constant_pool_countmagic;
	u2				minor_version
	u2				major_version
	u2				constant_pool_count
	cp_info			onstant_pool[constant_pool_count-1]
	u2				access_flags
	u2				this_class
	u2				super_class
	u2				interface_count
	u2				interface[interfaces_count]
	u2				fields_count
	field_info		fields[fields_count]
	u2				methods_count
	method_info		methods[methods_count]
	u2				attributes_count
	attribute_info	attributes[attributes_count]
}
*/
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse []byte classData to ClassFile struct
func Parse(classData []byte) (classfile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	classreader := &ClassReader{classData}
	classfile = &ClassFile{}
	classfile.read(classreader)
	return
}

func (mClassFile *ClassFile) read(reader *ClassReader) {
	mClassFile.readAndCheckMagic(reader)
	mClassFile.readAndCheckVersion(reader)
	mClassFile.constantPool = readConstantPool(reader)
	mClassFile.accessFlags = reader.readUint16()
	mClassFile.thisClass = reader.readUint16()
	mClassFile.superClass = reader.readUint16()
	mClassFile.interfaces = reader.readUint16s()
	mClassFile.fields = readMembers(reader, mClassFile.constantPool)
	mClassFile.methods = readMembers(reader, mClassFile.constantPool)
	mClassFile.attributes = readAttributes(reader, mClassFile.constantPool)
}

// MajorVersion getter
func (mClassFile *ClassFile) MajorVersion() uint16 {
	return mClassFile.majorVersion
}

//MinorVersion getter
func (mClassFile *ClassFile) MinorVersion() uint16 {
	return mClassFile.minorVersion
}

//ConstantPool getter
func (mClassFile *ClassFile) ConstantPool() ConstantPool {
	return mClassFile.constantPool
}

//AccessFlags getter
func (mClassFile *ClassFile) AccessFlags() uint16 {
	return mClassFile.accessFlags
}

//Fields getter
func (mClassFile *ClassFile) Fields() []*MemberInfo {
	return mClassFile.fields
}

//Methods getter
func (mClassFile *ClassFile) Methods() []*MemberInfo {
	return mClassFile.methods
}

// ClassName getter
func (mClassFile *ClassFile) ClassName() string {
	return mClassFile.constantPool.getClassName(mClassFile.thisClass)
}

// SuperClassName getter
func (mClassFile *ClassFile) SuperClassName() string {
	if mClassFile.superClass > 0 {
		return mClassFile.constantPool.getClassName(mClassFile.superClass)
	}
	return ""
}

// InterfaceNames getter
func (mClassFile *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(mClassFile.interfaces))
	for i, cpIndex := range mClassFile.interfaces {
		interfaceNames[i] = mClassFile.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (mClassFile *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (mClassFile *ClassFile) readAndCheckVersion(reader *ClassReader) {
	mClassFile.minorVersion = reader.readUint16()
	mClassFile.majorVersion = reader.readUint16()
	switch mClassFile.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52: // my jdk9 major version 52
		if mClassFile.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (mClassFile *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range mClassFile.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
