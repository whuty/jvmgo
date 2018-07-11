package heap

import "jvmgo/ch10/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (mClassMember *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	mClassMember.accessFlags = memberInfo.AccessFlags()
	mClassMember.name = memberInfo.Name()
	mClassMember.descriptor = memberInfo.Descriptor()
}

func (mClassMember *ClassMember) IsPublic() bool {
	return 0 != mClassMember.accessFlags&ACC_PUBLIC
}
func (mClassMember *ClassMember) IsPrivate() bool {
	return 0 != mClassMember.accessFlags&ACC_PRIVATE
}
func (mClassMember *ClassMember) IsProtected() bool {
	return 0 != mClassMember.accessFlags&ACC_PROTECTED
}
func (mClassMember *ClassMember) IsStatic() bool {
	return 0 != mClassMember.accessFlags&ACC_STATIC
}
func (mClassMember *ClassMember) IsFinal() bool {
	return 0 != mClassMember.accessFlags&ACC_FINAL
}
func (mClassMember *ClassMember) IsSynthetic() bool {
	return 0 != mClassMember.accessFlags&ACC_SYNTHETIC
}

// getters
func (mClassMember *ClassMember) Name() string {
	return mClassMember.name
}
func (mClassMember *ClassMember) Descriptor() string {
	return mClassMember.descriptor
}
func (mClassMember *ClassMember) Class() *Class {
	return mClassMember.class
}

// jvms 5.4.4
func (mClassMember *ClassMember) isAccessibleTo(d *Class) bool {
	if mClassMember.IsPublic() {
		return true
	}
	c := mClassMember.class
	if mClassMember.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	if !mClassMember.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
