package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/rtda"
	"jvmgo/ch10/rtda/heap"
)

func init() {
	native.Register("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;",
		getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;",
		getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0", "(Ljava/lang/Class;)Z",
		desiredAssertionStatus0)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
// 这里不讨论断言,将false推入堆栈
func desiredAssertionStatus0(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushBoolean(false)
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}
