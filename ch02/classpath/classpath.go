package classpath

import (
	"os"
	"path/filepath"
)

// Classpath struct
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// Parse use -Xjre to parse boot and ext class path
//         use -classpath/-cp to parse user class path
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
func (mClassPath *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	mClassPath.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	mClassPath.extClasspath = newWildcardEntry(jreExtPath)
}
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func (mClassPath *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	mClassPath.userClasspath = newEntry(cpOption)
}

// ReadClass read class file from bootClasspath, extClasspath, userClasspath, in order
// the className passed to ReadClass not include .class
func (mClassPath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := mClassPath.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := mClassPath.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return mClassPath.userClasspath.readClass(className)
}
func (mClassPath *Classpath) String() string {
	return mClassPath.userClasspath.String()
}
