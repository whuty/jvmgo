package rtda

import "jvmgo/ch08/rtda/heap"

// Frame of stack , lower, localVars, operandStack, thread, method, nextPC
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int // the next instruction after the call
}

// NewFrame consturctor
// maxLocals, maxStack have been defined by the Code attribute in method_info
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// LocalVars getter
func (mFrame *Frame) LocalVars() LocalVars {
	return mFrame.localVars
}

// OperandStack getter
func (mFrame *Frame) OperandStack() *OperandStack {
	return mFrame.operandStack
}

// Thread getter
func (mFrame *Frame) Thread() *Thread {
	return mFrame.thread
}

func (mFrame *Frame) Method() *heap.Method {
	return mFrame.method
}

// NextPC getter
func (mFrame *Frame) NextPC() int {
	return mFrame.nextPC
}

// SetNextPC setter
func (mFrame *Frame) SetNextPC(nextPC int) {
	mFrame.nextPC = nextPC
}

func (mFrame *Frame) RevertNextPC() {
	mFrame.nextPC = mFrame.thread.pc
}
