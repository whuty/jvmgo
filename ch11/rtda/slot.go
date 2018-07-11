package rtda

import "jvmgo/ch11/rtda/heap"

// Slot in localVals table, according to jvm rules, slot can contain at least one int or one ref
// we define a struct, contain one int and one ref
type Slot struct {
	num int32
	ref *heap.Object
}
