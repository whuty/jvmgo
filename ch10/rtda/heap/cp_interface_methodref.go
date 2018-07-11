package heap

import "jvmgo/ch10/classfile"

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
	d := mInterfaceMethodRef.constantpool.class
	c := mInterfaceMethodRef.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, mInterfaceMethodRef.name, mInterfaceMethodRef.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	mInterfaceMethodRef.method = method
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
