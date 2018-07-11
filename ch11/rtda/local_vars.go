package rtda

import (
	"jvmgo/ch11/rtda/heap"
	"math"
)

// LocalVars table
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// 操作局部变量表和操作数栈的指令都是隐含类型信息的
// 我们没有真的实现boolean, byte, short, char 的存取方法,这些类型都能转换成int来处理

// SetInt setter
func (mLocalVars LocalVars) SetInt(index uint, val int32) {
	mLocalVars[index].num = val
}

// GetInt getter
func (mLocalVars LocalVars) GetInt(index uint) int32 {
	return mLocalVars[index].num
}

// SetFloat setter
func (mLocalVars LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	mLocalVars[index].num = int32(bits)
}

// GetFloat getter
func (mLocalVars LocalVars) GetFloat(index uint) float32 {
	bits := uint32(mLocalVars[index].num)
	return math.Float32frombits(bits)
}

// SetLong setter
// long consumes two slots
func (mLocalVars LocalVars) SetLong(index uint, val int64) {
	mLocalVars[index].num = int32(val)
	mLocalVars[index+1].num = int32(val >> 32)
}

// GetLong getter
func (mLocalVars LocalVars) GetLong(index uint) int64 {
	low := uint32(mLocalVars[index].num)
	high := uint32(mLocalVars[index+1].num)
	return int64(high)<<32 | int64(low)
}

// SetDouble setter
// double consumes two slots
func (mLocalVars LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	mLocalVars.SetLong(index, int64(bits))
}

// GetDouble getter
func (mLocalVars LocalVars) GetDouble(index uint) float64 {
	bits := uint64(mLocalVars.GetLong(index))
	return math.Float64frombits(bits)
}

// SetRef setter
func (mLocalVars LocalVars) SetRef(index uint, ref *heap.Object) {
	mLocalVars[index].ref = ref
}

// GetRef getter
func (mLocalVars LocalVars) GetRef(index uint) *heap.Object {
	return mLocalVars[index].ref
}

func (mLocalVars LocalVars) SetSlot(index uint, slot Slot) {
	mLocalVars[index] = slot
}

func (mLocalVars LocalVars) GetThis() *heap.Object {
	return mLocalVars.GetRef(0)
}
func (mLocalVars LocalVars) GetBoolean(index uint) bool {
	return mLocalVars.GetInt(index) == 1
}
