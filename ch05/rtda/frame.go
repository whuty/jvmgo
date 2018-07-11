package rtda

// Frame of stack , lower, localVars, operandStack, thread, nextPC
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int // the next instruction after the call
}

// NewFrame consturctor
// maxLocals, maxStack have been defined by the Code attribute in method_info
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
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

// NextPC getter
func (mFrame *Frame) NextPC() int {
	return mFrame.nextPC
}

// SetNextPC setter
func (mFrame *Frame) SetNextPC(nextPC int) {
	mFrame.nextPC = nextPC
}
