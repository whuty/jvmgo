package heap

// symbolic reference
type SymRef struct {
	constantpool *ConstantPool
	className    string
	class        *Class
}

func (mSymRef *SymRef) ResolvedClass() *Class {
	if mSymRef.class == nil {
		mSymRef.resolveClassRef()
	}
	return mSymRef.class
}

// jvms8 5.4.3.1
func (mSymRef *SymRef) resolveClassRef() {
	d := mSymRef.constantpool.class
	c := d.loader.LoadClass(mSymRef.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	mSymRef.class = c
}
