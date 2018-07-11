package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry implenments Entry
type ZipEntry struct {
	absDir string
	zipRC  *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir, nil}
}
func (mZipEntry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if mZipEntry.zipRC == nil {
		err := mZipEntry.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := mZipEntry.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, mZipEntry, err
}

func (mZipEntry *ZipEntry) openJar() error {
	r, err := zip.OpenReader(mZipEntry.absDir)
	if err == nil {
		mZipEntry.zipRC = r
	}
	return err
}

func (mZipEntry *ZipEntry) findClass(className string) *zip.File {
	for _, f := range mZipEntry.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (mZipEntry *ZipEntry) String() string {
	return mZipEntry.absDir
}
