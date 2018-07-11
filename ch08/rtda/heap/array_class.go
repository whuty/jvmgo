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
		return &Object{mClass, make([]int8, count)}
	case "[B":
		return &Object{mClass, make([]int8, count)}
	case "[C":
		return &Object{mClass, make([]uint16, count)}
	case "[S":
		return &Object{mClass, make([]int16, count)}
	case "[I":
		return &Object{mClass, make([]int32, count)}
	case "[J":
		return &Object{mClass, make([]int64, count)}
	case "[F":
		return &Object{mClass, make([]float32, count)}
	case "[D":
		return &Object{mClass, make([]float64, count)}
	default:
		return &Object{mClass, make([]*Object, count)}
	}
}
