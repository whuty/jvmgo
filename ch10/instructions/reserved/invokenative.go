package reserved

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/native"
	"jvmgo/ch10/rtda"

	_ "jvmgo/ch10/native/java/lang" // import for side effect
	_ "jvmgo/ch10/native/sun/misc"
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
