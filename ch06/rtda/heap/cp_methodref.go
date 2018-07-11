package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.constantpool = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (mMethodref *MethodRef) ResolvedMethod() *Method {
	if mMethodref.method == nil {
		mMethodref.resolveMethodRef()
	}
	return mMethodref.method
}

// jvms8 5.4.3.3
func (mMethodref *MethodRef) resolveMethodRef() {
	//class := mMethodref.Class()
	// todo
}
