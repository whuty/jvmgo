package rtda

import "math"

// OperandStack 操作数栈
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// PushInt method
func (mOperandStack *OperandStack) PushInt(val int32) {
	mOperandStack.slots[mOperandStack.size].num = val
	mOperandStack.size++
}

// PopInt method
func (mOperandStack *OperandStack) PopInt() int32 {
	mOperandStack.size--
	return mOperandStack.slots[mOperandStack.size].num
}

// PushFloat method
func (mOperandStack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	mOperandStack.slots[mOperandStack.size].num = int32(bits)
	mOperandStack.size++
}

// PopFloat method
func (mOperandStack *OperandStack) PopFloat() float32 {
	mOperandStack.size--
	bits := uint32(mOperandStack.slots[mOperandStack.size].num)
	return math.Float32frombits(bits)
}

// PushLong method
// long consumes two slots
func (mOperandStack *OperandStack) PushLong(val int64) {
	mOperandStack.slots[mOperandStack.size].num = int32(val)
	mOperandStack.slots[mOperandStack.size+1].num = int32(val >> 32)
	mOperandStack.size += 2
}

// PopLong method
func (mOperandStack *OperandStack) PopLong() int64 {
	mOperandStack.size -= 2
	low := uint32(mOperandStack.slots[mOperandStack.size].num)
	high := uint32(mOperandStack.slots[mOperandStack.size+1].num)
	return int64(high)<<32 | int64(low)
}

// PushDouble method
// double consumes two slots
func (mOperandStack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	mOperandStack.PushLong(int64(bits))
}

// PopDouble method
func (mOperandStack *OperandStack) PopDouble() float64 {
	bits := uint64(mOperandStack.PopLong())
	return math.Float64frombits(bits)
}

// PushRef method
func (mOperandStack *OperandStack) PushRef(ref *Object) {
	mOperandStack.slots[mOperandStack.size].ref = ref
	mOperandStack.size++
}

// PopRef method
func (mOperandStack *OperandStack) PopRef() *Object {
	mOperandStack.size--
	ref := mOperandStack.slots[mOperandStack.size].ref
	mOperandStack.slots[mOperandStack.size].ref = nil
	return ref
}
