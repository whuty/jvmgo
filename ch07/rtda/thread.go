package rtda

import "jvmgo/ch07/rtda/heap"

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

// PC getter
func (mThread *Thread) PC() int {
	return mThread.pc
}

// SetPC setter
func (mThread *Thread) SetPC(pc int) {
	mThread.pc = pc
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
func (mThread *Thread) TopFrame() *Frame {
	return mThread.stack.top()
}

func (mThread *Thread) IsStackEmpty() bool {
	return mThread.stack.isEmpty()
}

// NewFrame .
func (mThread *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(mThread, method)
}
