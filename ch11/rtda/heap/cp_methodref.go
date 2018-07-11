package heap

import "jvmgo/ch11/classfile"

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
	d := mMethodref.constantpool.class
	c := mMethodref.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(c, mMethodref.name, mMethodref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	mMethodref.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
