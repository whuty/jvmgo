package heap

import (
	"jvmgo/ch10/classfile"
)

type Method struct {
	ClassMember
	maxStack        uint
	maxLocals       uint
	code            []byte
	exceptionTable  ExceptionTable
	lineNumberTable *classfile.LineNumberTableAttribute
	argSlotCount    uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfmethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfmethod)
	method.copyAttributes(cfmethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (mMethod *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		mMethod.maxStack = codeAttr.MaxStack()
		mMethod.maxLocals = codeAttr.MaxLocals()
		mMethod.code = codeAttr.Code()
		mMethod.lineNumberTable = codeAttr.LineNumberTableAttribute()
		mMethod.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			mMethod.class.constantPool)
	}
}

func (mMethod *Method) GetLineNumber(pc int) int {
	if mMethod.IsNative() {
		return -2
	}
	if mMethod.lineNumberTable == nil {
		return -1
	}
	return mMethod.lineNumberTable.GetLineNumber(pc)
}

func (mMethod *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := mMethod.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (mMethod *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		mMethod.argSlotCount++
		if paramType == "J" || paramType == "D" {
			mMethod.argSlotCount++
		}
	}
	if !mMethod.IsStatic() {
		mMethod.argSlotCount++
	}
}

func (mMethod *Method) injectCodeAttribute(returnType string) {
	mMethod.maxStack = 4
	mMethod.maxLocals = mMethod.argSlotCount
	switch returnType[0] {
	case 'V':
		mMethod.code = []byte{0xfe, 0xb1} // return
	case 'D':
		mMethod.code = []byte{0xfe, 0xaf} //dreturn
	case 'F':
		mMethod.code = []byte{0xfe, 0xae} //freturn
	case 'J':
		mMethod.code = []byte{0xfe, 0xad} //lreturn
	case 'L', '[':
		mMethod.code = []byte{0xfe, 0xb0} //areturn
	default:
		mMethod.code = []byte{0xfe, 0xac} //ireturn
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
