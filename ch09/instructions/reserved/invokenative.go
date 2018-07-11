package reserved

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"

	_ "jvmgo/ch09/native/java/lang" // import for side effect
	_ "jvmgo/ch09/native/sun/misc"
)

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
