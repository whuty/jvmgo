package rtda

// Thread , define pc, stack
type Thread struct {
	pc    int
	stack *Stack
}

// NewThread , init stack size is 1024
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// PushFrame , call stack push method
func (mThread *Thread) PushFrame(frame *Frame) {
	mThread.stack.push(frame)
}

// PopFrame , call stack pop method
func (mThread *Thread) PopFrame() *Frame {
	return mThread.stack.pop()
}

// CurrentFrame , return current frame
func (mThread *Thread) CurrentFrame() *Frame {
	return mThread.stack.top()
}
