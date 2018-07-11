package heap

import "jvmgo/ch10/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.constantpool = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (mFieldRef *FieldRef) ResolvedField() *Field {
	if mFieldRef.field == nil {
		mFieldRef.resolveFieldRef()
	}
	return mFieldRef.field
}

// jvms 5.4.3.2
func (mFieldRef *FieldRef) resolveFieldRef() {
	d := mFieldRef.constantpool.class
	c := mFieldRef.ResolvedClass()
	field := lookupField(c, mFieldRef.name, mFieldRef.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	mFieldRef.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
