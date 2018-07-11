package main

import (
	"fmt"
	"jvmgo/jvmgo/classpath"
	"jvmgo/jvmgo/instructions/base"
	"jvmgo/jvmgo/rtda"
	"jvmgo/jvmgo/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (mJVM *JVM) start() {
	mJVM.initVM()
	mJVM.execMain()
}

func (mJVM *JVM) initVM() {
	vmClass := mJVM.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(mJVM.mainThread, vmClass)
	interpret(mJVM.mainThread, mJVM.cmd.verboseInstFlag)
}

func (mJVM *JVM) execMain() {
	className := strings.Replace(mJVM.cmd.class, ".", "/", -1)
	mainClass := mJVM.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", mJVM.cmd.class)
		return
	}

	argsArr := mJVM.createArgsArray()
	frame := mJVM.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	mJVM.mainThread.PushFrame(frame)
	interpret(mJVM.mainThread, mJVM.cmd.verboseInstFlag)
}

func (mJVM *JVM) createArgsArray() *heap.Object {
	stringClass := mJVM.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(mJVM.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range mJVM.cmd.args {
		jArgs[i] = heap.JString(mJVM.classLoader, arg)
	}
	return argsArr
}
