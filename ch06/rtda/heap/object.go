package heap

// Object 对象类型
type Object struct {
	class  *Class
	fields *Slots
}

func newObject(class *Class) *Object {
	slots := newSlots(class.instanceSlotCount)
	return &Object{
		class:  class,
		fields: &slots,
	}
}

// getters
func (mObj *Object) Class() *Class {
	return mObj.class
}
func (mObj *Object) Fields() *Slots {
	return mObj.fields
}

func (mObj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(mObj.class)
}
