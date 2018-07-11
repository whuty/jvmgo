package rtda

// Stack struct for jvm, we use linked list here
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (mStack *Stack) push(frame *Frame) {
	if mStack.size >= mStack.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if mStack._top != nil {
		frame.lower = mStack._top
	}

	mStack._top = frame
	mStack.size++
}

func (mStack *Stack) pop() *Frame {
	if mStack._top == nil {
		panic("jvm stack is empty!")
	}

	top := mStack._top
	mStack._top = top.lower
	top.lower = nil
	mStack.size--

	return top
}

func (mStack *Stack) top() *Frame {
	if mStack._top == nil {
		panic("jvm stack is empty!")
	}

	return mStack._top
}

func (mStack *Stack) isEmpty() bool {
	return mStack._top == nil
}
