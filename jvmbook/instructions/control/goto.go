package control

import "jvmgo/jvmbook/instructions/base"
import "jvmgo/jvmbook/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
