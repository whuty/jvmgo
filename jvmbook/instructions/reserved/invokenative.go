package reserved

import "jvmgo/jvmbook/instructions/base"
import "jvmgo/jvmbook/rtda"
import "jvmgo/jvmbook/native"
import _ "jvmgo/jvmbook/native/java/io"
import _ "jvmgo/jvmbook/native/java/lang"
import _ "jvmgo/jvmbook/native/java/security"
import _ "jvmgo/jvmbook/native/java/util/concurrent/atomic"
import _ "jvmgo/jvmbook/native/sun/io"
import _ "jvmgo/jvmbook/native/sun/misc"
import _ "jvmgo/jvmbook/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
