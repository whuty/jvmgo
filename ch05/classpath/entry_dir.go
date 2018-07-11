package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry implements Entry
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (mDirEntry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(mDirEntry.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, mDirEntry, err
}
func (mDirEntry *DirEntry) String() string {
	return mDirEntry.absDir
}
