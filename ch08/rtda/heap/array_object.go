package heap

func (mObj *Object) Bytes() []int8 {
	return mObj.data.([]int8)
}
func (mObj *Object) Shorts() []int16 {
	return mObj.data.([]int16)
}
func (mObj *Object) Ints() []int32 {
	return mObj.data.([]int32)
}
func (mObj *Object) Longs() []int64 {
	return mObj.data.([]int64)
}
func (mObj *Object) Chars() []uint16 {
	return mObj.data.([]uint16)
}
func (mObj *Object) Floats() []float32 {
	return mObj.data.([]float32)
}
func (mObj *Object) Doubles() []float64 {
	return mObj.data.([]float64)
}
func (mObj *Object) Refs() []*Object {
	return mObj.data.([]*Object)
}

func (mObj *Object) ArrayLength() int32 {
	switch mObj.data.(type) {
	case []int8:
		return int32(len(mObj.data.([]int8)))
	case []int16:
		return int32(len(mObj.data.([]int16)))
	case []int32:
		return int32(len(mObj.data.([]int32)))
	case []int64:
		return int32(len(mObj.data.([]int64)))
	case []uint16:
		return int32(len(mObj.data.([]uint16)))
	case []float32:
		return int32(len(mObj.data.([]float32)))
	case []float64:
		return int32(len(mObj.data.([]float64)))
	case []*Object:
		return int32(len(mObj.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
