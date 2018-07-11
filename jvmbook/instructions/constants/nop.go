package constants

import "jvmgo/jvmbook/instructions/base"
import "jvmgo/jvmbook/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
