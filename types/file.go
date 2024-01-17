package types

import (
	"fmt"
	"os"
	"sync"
)

type Files struct {
	Files []File
}

type File struct {
	os.FileInfo
	Lock    sync.Mutex
	Content []byte
}

func (f *Files) FindByName(fileName *string) (*File, error) {
	if len(f.Files) == 0 {
		return nil, fmt.Errorf("%s not found", *fileName)
	}
	for v := range f.Files {
		if f.Files[v].FileInfo.Name() == *fileName {
			return &f.Files[v], nil
		}
	}
	return nil, fmt.Errorf("%s not found", *fileName)
}

func (f *Files) DeleteByName(fileName *string) error {
	if len(f.Files) == 0 {
		return fmt.Errorf("%s not found", *fileName)
	}
	for v := range f.Files {
		if f.Files[v].FileInfo.Name() == *fileName {
			f.Files = append(f.Files[:v], f.Files[v+1:]...)
			break
		}
	}
	return nil
}

func (f *Files) New(fileName *string, data []byte) error {
	fileDescriptor, err := os.OpenFile(*fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("unable to create file descriptor: %s", err)
	}

	defer fileDescriptor.Close()

	_, err = fileDescriptor.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write data to file: %s", err)
	}

	var newFile File
	newFile.FileInfo, err = os.Stat(*fileName)
	if err != nil {
		return fmt.Errorf("unable to get info of created file: %s", err)
	}
	newFile.Content = data
	f.Files = append(f.Files, newFile)
	return nil
}

func (f *File) Modify(data []byte) error {
	f.Lock.Lock()
	fileDescriptor, err := os.OpenFile(f.FileInfo.Name(), os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("unable to open file: %s", err)
	}

	defer func() {
		fileDescriptor.Close()
		f.Lock.Unlock()
	}()
	_, err = fileDescriptor.Write(data)
	if err != nil {
		return fmt.Errorf("unable to write data to file: %s", err)
	}
	f.Content = data
	return nil
}

func (f *File) Delete() error {
	f.Lock.Lock()
	defer f.Lock.Unlock()
	err := os.Remove(f.FileInfo.Name())
	if err != nil {
		return err
	}
	return nil
}
