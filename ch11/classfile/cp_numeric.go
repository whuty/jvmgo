package classfile

import (
	"math"
)

/*
ConstantIntegerInfo {
	u1 tag;
	u4 bytes;

}
actually in this implement the types smaller than int
boolean, byte, short, char are in Integer Info as well
*/
type ConstantIntegerInfo struct {
	val int32
}

func (mConstantIntegerInfo *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	mConstantIntegerInfo.val = int32(bytes)
}

//Value getter
func (mConstantIntegerInfo *ConstantIntegerInfo) Value() int32 {
	return mConstantIntegerInfo.val
}

/*
ConstantFloatInfo {
	u1 tag
	u4 bytes
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (mConstantFloatInfo *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	mConstantFloatInfo.val = math.Float32frombits(bytes)
}

//Value getter
func (mConstantFloatInfo *ConstantFloatInfo) Value() float32 {
	return mConstantFloatInfo.val
}

/*
ConstantLongInfo {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (mConstantLongInfo *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	mConstantLongInfo.val = int64(bytes)
}

//Value getter
func (mConstantLongInfo *ConstantLongInfo) Value() int64 {
	return mConstantLongInfo.val
}

/*
ConstantDoubleInfo {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (mConstantDoubleInfo *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	mConstantDoubleInfo.val = math.Float64frombits(bytes)
}

// Value getter
func (mConstantDoubleInfo *ConstantDoubleInfo) Value() float64 {
	return mConstantDoubleInfo.val
}
