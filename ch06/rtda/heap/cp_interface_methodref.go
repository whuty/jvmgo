package heap

import "jvmgo/ch06/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.constantpool = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (mInterfaceMethodRef *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if mInterfaceMethodRef.method == nil {
		mInterfaceMethodRef.resolveInterfaceMethodRef()
	}
	return mInterfaceMethodRef.method
}

// jvms8 5.4.3.4
func (mInterfaceMethodRef *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := mInterfaceMethodRef.ResolveClass()
	// todo
}
