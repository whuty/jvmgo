package heap

import (
	"jvmgo/jvmgo/classfile"
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
	sourceFile        string
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
	jClass            *Object
	initStarted       bool //
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
	class.sourceFile = getSourceFile(cf)
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
func (mClass *Class) Name() string {
	return mClass.name
}
func (mClass *Class) AccessFlags() uint16 {
	return mClass.accessFlags
}
func (mClass *Class) ConstantPool() *ConstantPool {
	return mClass.constantPool
}
func (mClass *Class) StaticVars() *Slots {
	return mClass.staticVars
}
func (mClass *Class) SuperClass() *Class {
	return mClass.superClass
}
func (mClass *Class) Loader() *ClassLoader {
	return mClass.loader
}
func (mClass *Class) Interfaces() []*Class {
	return mClass.interfaces
}
func (mClass *Class) JClass() *Object {
	return mClass.jClass
}
func (mClass *Class) JavaName() string {
	return strings.Replace(mClass.name, "/", ".", -1)
}
func (mClass *Class) SourceFile() string {
	return mClass.sourceFile
}
func (mClass *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[mClass.name]
	return ok
}

// jvms 5.4.4
func (mClass *Class) isAccessibleTo(other *Class) bool {
	return mClass.IsPublic() ||
		mClass.GetPackageName() == other.GetPackageName()
}

func (mClass *Class) GetPackageName() string {
	if i := strings.LastIndex(mClass.name, "/"); i >= 0 {
		return mClass.name[:i]
	}
	return ""
}

func (mClass *Class) GetMainMethod() *Method {
	return mClass.GetStaticMethod("main", "([Ljava/lang/String;)V")
}
func (mClass *Class) GetClinitMethod() *Method {
	return mClass.GetStaticMethod("<clinit>", "()V")
}
func (mClass *Class) GetStaticMethod(name, descriptor string) *Method {
	for _, method := range mClass.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}
func (mClass *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := mClass; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}
func (mClass *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := mClass; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}

func (mClass *Class) isJlObject() bool {
	return mClass.name == "java/lang/Object"
}
func (mClass *Class) isJlCloneable() bool {
	return mClass.name == "java/lang/Cloneable"
}
func (mClass *Class) isJioSerializable() bool {
	return mClass.name == "java/io/Serializable"
}

func (mClass *Class) NewObject() *Object {
	return newObject(mClass)
}

func (mClass *Class) InitStarted() bool {
	return mClass.initStarted
}
func (mClass *Class) StartInit() {
	mClass.initStarted = true
}

func (mClass *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(mClass.name)
	return mClass.loader.LoadClass(arrayClassName)
}
func (mClass *Class) GetInstanceMethod(name, descriptor string) *Method {
	return mClass.getMethod(name, descriptor, false)
}

func (mClass *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := mClass.getField(fieldName, fieldDescriptor, true)
	return mClass.staticVars.GetRef(field.slotId)
}
func (mClass *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := mClass.getField(fieldName, fieldDescriptor, true)
	mClass.staticVars.SetRef(field.slotId, ref)
}

func (mClass *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(mClass.fields))
		for _, field := range mClass.fields {
			if field.IsPublic() {
				publicFields = append(publicFields, field)
			}
		}
		return publicFields
	} else {
		return mClass.fields
	}
}

func (mClass *Class) GetConstructor(descriptor string) *Method {
	return mClass.GetInstanceMethod("<init>", descriptor)
}

func (mClass *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(mClass.methods))
	for _, method := range mClass.methods {
		if method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				constructors = append(constructors, method)
			}
		}
	}
	return constructors
}

func (mClass *Class) GetMethods(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(mClass.methods))
	for _, method := range mClass.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}
