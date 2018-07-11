package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		slots := make([]Slot, slotCount)
		return slots
	}
	return nil
}

func (mSlots Slots) SetInt(index uint, val int32) {
	mSlots[index].num = val
}
func (mSlots Slots) GetInt(index uint) int32 {
	return mSlots[index].num
}

func (mSlots Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	mSlots[index].num = int32(bits)
}
func (mSlots Slots) GetFloat(index uint) float32 {
	bits := uint32(mSlots[index].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (mSlots Slots) SetLong(index uint, val int64) {
	mSlots[index].num = int32(val)
	mSlots[index+1].num = int32(val >> 32)
}
func (mSlots Slots) GetLong(index uint) int64 {
	low := uint32(mSlots[index].num)
	high := uint32(mSlots[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (mSlots Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	mSlots.SetLong(index, int64(bits))
}
func (mSlots Slots) GetDouble(index uint) float64 {
	bits := uint64(mSlots.GetLong(index))
	return math.Float64frombits(bits)
}

func (mSlots Slots) SetRef(index uint, ref *Object) {
	mSlots[index].ref = ref
}
func (mSlots Slots) GetRef(index uint) *Object {
	return mSlots[index].ref
}
