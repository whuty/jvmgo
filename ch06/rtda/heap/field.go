package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfFields := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfFields)
	}
	return fields
}
func (mField *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		mField.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (mField *Field) IsVolatile() bool {
	return 0 != mField.accessFlags&ACC_VOLATILE
}
func (mField *Field) IsTransient() bool {
	return 0 != mField.accessFlags&ACC_TRANSIENT
}
func (mField *Field) IsEnum() bool {
	return 0 != mField.accessFlags&ACC_ENUM
}

func (mField *Field) ConstValueIndex() uint {
	return mField.constValueIndex
}
func (mField *Field) SlotId() uint {
	return mField.slotId
}
func (mField *Field) isLongOrDouble() bool {
	return mField.descriptor == "J" || mField.descriptor == "D"
}
