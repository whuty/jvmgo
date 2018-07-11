package atomic

import "jvmgo/jvmbook/native"
import "jvmgo/jvmbook/rtda"

func init() {
	native.Register("java/util/concurrent/atomic/AtomicLong", "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}
