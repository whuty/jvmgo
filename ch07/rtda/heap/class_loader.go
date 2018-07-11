package heap

import (
	"fmt"
	"jvmgo/ch07/classfile"
	"jvmgo/ch07/classpath"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	classpath   *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		classpath:   cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
}

func (mClassLoader *ClassLoader) LoadClass(name string) *Class {
	if class, ok := mClassLoader.classMap[name]; ok {
		// already loaded
		return class
	}

	return mClassLoader.loadNonArrayClass(name)
}

func (mClassLoader *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := mClassLoader.readClass(name)
	class := mClassLoader.defineClass(data)
	link(class)
	if mClassLoader.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (mClassLoader *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := mClassLoader.classpath.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
func (mClassLoader *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = mClassLoader
	resolveSuperClass(class)
	resolveInterfaces(class)
	mClassLoader.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// jvms 5.4.2
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	slots := newSlots(class.staticSlotCount)
	class.staticVars = &slots
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
