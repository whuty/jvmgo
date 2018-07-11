package classfile

import (
	"encoding/binary"
)

/*
Class file{
	u4				magic;
	u2				minor_version
	u2				marjor_version
	u2				constant_pool_countmagic;
	u2				minor_version
	u2				marjor_version
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

/*
	go			java		detail
	int8		byte		8 bit signed int
	uint8(byte)	--			8 bit unsigned int
	int16		short		16 bit signed int
	uint16		char
	int32(rune)	int
	uint32		--
	int64		long
	uint64		--
	float32		float		IEEE-754 float
	float64		double		IEEE-754 double
*/

// ClassReader read class file
type ClassReader struct {
	data []byte
}

// read u1
func (mClassReader *ClassReader) readUint8() uint8 {
	val := mClassReader.data[0]
	mClassReader.data = mClassReader.data[1:]
	return val
}

// read u2
func (mClassReader *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(mClassReader.data)
	mClassReader.data = mClassReader.data[2:]
	return val
}

// read u4
func (mClassReader *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(mClassReader.data)
	mClassReader.data = mClassReader.data[4:]
	return val
}

// read u8
// note, u8 is not defined in jvm
func (mClassReader *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(mClassReader.data)
	mClassReader.data = mClassReader.data[8:]
	return val
}

// read uint16(u2) table
// the size of table is defined by first uint16
func (mClassReader *ClassReader) readUint16s() []uint16 {
	n := mClassReader.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = mClassReader.readUint16()
	}
	return s
}

// read bytes
func (mClassReader *ClassReader) readBytes(n uint32) []byte {
	bytes := mClassReader.data[:n]
	mClassReader.data = mClassReader.data[n:]
	return bytes
}
