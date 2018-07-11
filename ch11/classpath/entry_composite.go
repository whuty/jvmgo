package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry array of Entry
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
func (mCompositeEntry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range mCompositeEntry {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (mCompositeEntry CompositeEntry) String() string {
	strs := make([]string, len(mCompositeEntry))
	for i, entry := range mCompositeEntry {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
