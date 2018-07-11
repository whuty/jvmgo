package heap

import "jvmgo/ch08/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}
func (mMethod *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		mMethod.maxStack = codeAttr.MaxStack()
		mMethod.maxLocals = codeAttr.MaxLocals()
		mMethod.code = codeAttr.Code()
	}
}

func (mMethod *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(mMethod.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		mMethod.argSlotCount++
		if paramType == "J" || paramType == "D" {
			mMethod.argSlotCount++
		}
	}
	if !mMethod.IsStatic() {
		mMethod.argSlotCount++
	}
}

func (mMethod *Method) IsSynchronized() bool {
	return 0 != mMethod.accessFlags&ACC_SYNCHRONIZED
}
func (mMethod *Method) IsBridge() bool {
	return 0 != mMethod.accessFlags&ACC_BRIDGE
}
func (mMethod *Method) IsVarargs() bool {
	return 0 != mMethod.accessFlags&ACC_VARARGS
}
func (mMethod *Method) IsNative() bool {
	return 0 != mMethod.accessFlags&ACC_NATIVE
}
func (mMethod *Method) IsAbstract() bool {
	return 0 != mMethod.accessFlags&ACC_ABSTRACT
}
func (mMethod *Method) IsStrict() bool {
	return 0 != mMethod.accessFlags&ACC_STRICT
}

// getters
func (mMethod *Method) MaxStack() uint {
	return mMethod.maxStack
}
func (mMethod *Method) MaxLocals() uint {
	return mMethod.maxLocals
}
func (mMethod *Method) Code() []byte {
	return mMethod.code
}
func (mMethod *Method) ArgSlotCount() uint {
	return mMethod.argSlotCount
}
