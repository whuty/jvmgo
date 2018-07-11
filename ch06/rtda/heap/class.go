package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (mClass *Class) IsPublic() bool {
	return 0 != mClass.accessFlags&ACC_PUBLIC
}
func (mClass *Class) IsFinal() bool {
	return 0 != mClass.accessFlags&ACC_FINAL
}
func (mClass *Class) IsSuper() bool {
	return 0 != mClass.accessFlags&ACC_SUPER
}
func (mClass *Class) IsInterface() bool {
	return 0 != mClass.accessFlags&ACC_INTERFACE
}
func (mClass *Class) IsAbstract() bool {
	return 0 != mClass.accessFlags&ACC_ABSTRACT
}
func (mClass *Class) IsSynthetic() bool {
	return 0 != mClass.accessFlags&ACC_SYNTHETIC
}
func (mClass *Class) IsAnnotation() bool {
	return 0 != mClass.accessFlags&ACC_ANNOTATION
}
func (mClass *Class) IsEnum() bool {
	return 0 != mClass.accessFlags&ACC_ENUM
}

// getters
func (mClass *Class) ConstantPool() *ConstantPool {
	return mClass.constantPool
}
func (mClass *Class) StaticVars() *Slots {
	return mClass.staticVars
}

// jvms 5.4.4
func (mClass *Class) isAccessibleTo(other *Class) bool {
	return mClass.IsPublic() ||
		mClass.getPackageName() == other.getPackageName()
}

func (mClass *Class) getPackageName() string {
	if i := strings.LastIndex(mClass.name, "/"); i >= 0 {
		return mClass.name[:i]
	}
	return ""
}

func (mClass *Class) GetMainMethod() *Method {
	return mClass.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (mClass *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range mClass.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (mClass *Class) NewObject() *Object {
	return newObject(mClass)
}
