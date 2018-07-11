package lang

import (
	"jvmgo/jvmgo/native"
	"jvmgo/jvmgo/rtda"
	"math"
)

func init() {
	native.Register("java/lang/StrictMath", "pow", "(DD)D", pow)
}

// public static native double pow(double a, double b);
// (DD)D;
func pow(frame *rtda.Frame) {
	x := frame.LocalVars().GetDouble(0)
	y := frame.LocalVars().GetDouble(2)
	res := math.Pow(x, y)
	frame.OperandStack().PushDouble(res)
}
