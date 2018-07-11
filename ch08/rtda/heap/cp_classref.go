package heap

import "jvmgo/ch08/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.constantpool = cp
	ref.className = classInfo.Name()
	return ref
}
