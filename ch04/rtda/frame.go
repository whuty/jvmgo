package rtda

// Frame of stack , lower, localVars, operandStack, it is simple now
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

// NewFrame consturctor
// maxLocals, maxStack have been defined by the Code attribute in method_info
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
