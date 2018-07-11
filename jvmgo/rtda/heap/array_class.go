package heap

func (mClass *Class) IsArray() bool {
	return mClass.name[0] == '['
}

func (mClass *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(mClass.name)
	return mClass.loader.LoadClass(componentClassName)
}

func (mClass *Class) NewArray(count uint) *Object {
	if !mClass.IsArray() {
		panic("Not array class: " + mClass.name)
	}
	switch mClass.Name() {
	case "[Z":
		return &Object{mClass, make([]int8, count), nil}
	case "[B":
		return &Object{mClass, make([]int8, count), nil}
	case "[C":
		return &Object{mClass, make([]uint16, count), nil}
	case "[S":
		return &Object{mClass, make([]int16, count), nil}
	case "[I":
		return &Object{mClass, make([]int32, count), nil}
	case "[J":
		return &Object{mClass, make([]int64, count), nil}
	case "[F":
		return &Object{mClass, make([]float32, count), nil}
	case "[D":
		return &Object{mClass, make([]float64, count), nil}
	default:
		return &Object{mClass, make([]*Object, count), nil}
	}
}

func NewByteArray(loader *ClassLoader, bytes []int8) *Object {
	return &Object{loader.LoadClass("[B"), bytes, nil}
}
