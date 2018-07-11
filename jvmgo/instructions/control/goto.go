package control

import (
	"jvmgo/jvmgo/instructions/base"
	"jvmgo/jvmgo/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
