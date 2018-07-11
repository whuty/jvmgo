package rtda

import "jvmgo/jvmbook/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
