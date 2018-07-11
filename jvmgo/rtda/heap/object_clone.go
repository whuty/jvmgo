package heap

func (mObject *Object) Clone() *Object {
	return &Object{
		class: mObject.class,
		data:  mObject.cloneData(),
	}
}

func (mObject *Object) cloneData() interface{} {
	switch mObject.data.(type) {
	case []int8:
		elements := mObject.data.([]int8)
		elements2 := make([]int8, len(elements))
		copy(elements2, elements)
		return elements2
	case []int16:
		elements := mObject.data.([]int16)
		elements2 := make([]int16, len(elements))
		copy(elements2, elements)
		return elements2
	case []uint16:
		elements := mObject.data.([]uint16)
		elements2 := make([]uint16, len(elements))
		copy(elements2, elements)
		return elements2
	case []int32:
		elements := mObject.data.([]int32)
		elements2 := make([]int32, len(elements))
		copy(elements2, elements)
		return elements2
	case []int64:
		elements := mObject.data.([]int64)
		elements2 := make([]int64, len(elements))
		copy(elements2, elements)
		return elements2
	case []float32:
		elements := mObject.data.([]float32)
		elements2 := make([]float32, len(elements))
		copy(elements2, elements)
		return elements2
	case []float64:
		elements := mObject.data.([]float64)
		elements2 := make([]float64, len(elements))
		copy(elements2, elements)
		return elements2
	case []*Object:
		elements := mObject.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default: // []Slot
		slots := mObject.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}
