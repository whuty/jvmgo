package heap

// Object 对象类型
type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (mObj *Object) Class() *Class {
	return mObj.class
}
func (mObj *Object) Data() interface{} {
	return mObj.data
}
func (mObj *Object) Fields() Slots {
	return mObj.data.(Slots)
}
func (mObj *Object) Extra() interface{} {
	return mObj.extra
}
func (mObj *Object) SetExtra(extra interface{}) {
	mObj.extra = extra
}
func (mObj *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(mObj.class)
}

// reflection
func (mObj *Object) GetRefVar(name, descriptor string) *Object {
	field := mObj.class.getField(name, descriptor, false)
	slots := mObj.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (mObj *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := mObj.class.getField(name, descriptor, false)
	slots := mObj.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (mObj *Object) SetIntVar(name, descriptor string, val int32) {
	field := mObj.class.getField(name, descriptor, false)
	slots := mObj.data.(Slots)
	slots.SetInt(field.slotId, val)
}
func (mObj *Object) GetIntVar(name, descriptor string) int32 {
	field := mObj.class.getField(name, descriptor, false)
	slots := mObj.data.(Slots)
	return slots.GetInt(field.slotId)
}
