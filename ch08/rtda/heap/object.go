package heap

// Object 对象类型
type Object struct {
	class *Class
	data  interface{}
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
func (mObj *Object) Fields() Slots {
	return mObj.data.(Slots)
}

func (mObj *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(mObj.class)
}

// reflection
func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
