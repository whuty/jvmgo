package heap

import "jvmgo/ch09/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (mMemberRef *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	mMemberRef.className = refInfo.ClassName()
	mMemberRef.name, mMemberRef.descriptor = refInfo.NameAndDescriptor()
}

func (mMemberRef *MemberRef) Name() string {
	return mMemberRef.name
}
func (mMemberRef *MemberRef) Descriptor() string {
	return mMemberRef.descriptor
}
