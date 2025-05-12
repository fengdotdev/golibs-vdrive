package vdrive

import (
	"encoding/json"
	"os"
	"path"
)

type route = string

type identifier = string

const (
	INDEXFILE = "index-VDRIVE.json"
)

type VDrive struct {
	workingDir string
	index      map[route]identifier // map[relativePath]id
}

func NewVDrive() *VDrive {
	wd, err := os.Getwd()
	if err != nil {
		panic("failed to get working directory")
	}
	return &VDrive{
		workingDir: wd,
		index:      make(map[string]string),
	}
}

func (v *VDrive) fullpath(relativePath string) (string, error) {

}

func (v *VDrive) LoadIndex() error {
	fullpath := path.Join(v.workingDir, INDEXFILE)

	file, err := os.Open(fullpath)
	if err != nil {
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&v.index)
	if err != nil {
		return err
	}

	// Check if the index is empty
	if len(v.index) == 0 {
		return os.ErrNotExist
	}

	return nil

}

func (v *VDrive) SaveIndex() error {
	fullpath := path.Join(v.workingDir, INDEXFILE)

	data, err := json.Marshal(v.index)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (v *VDrive) CreateFile(rawPath string, data []byte) error {

	id := GenerateID()
	relativePath := RelativePath(rawPath)

	// Check if the file already exists
	if _, exists := v.index[relativePath]; exists {
		return ErrFileExists
	}

	fullpath := path.Join(v.workingDir, id)

	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	v.index[relativePath] = fullpath

	return nil
}

func (v *VDrive) ReadFile(relativePath string) ([]byte, error) {
	relativePath = RelativePath(relativePath)

	fullpath, exists := v.index[relativePath]
	if !exists {
		return nil, ErrNotExist
	}

	file, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := os.ReadFile(fullpath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (v *VDrive) DeleteFile(relativePath string) error {
	relativePath = RelativePath(relativePath)

	fullpath, exists := v.index[relativePath]
	if !exists {
		return ErrNotExist
	}

	err := os.Remove(fullpath)
	if err != nil {
		return err
	}

	delete(v.index, relativePath)

	return nil
}

func (v *VDrive) UpdateFile(relativePath string, data []byte) error {
	relativePath = RelativePath(relativePath)

	fullpath, exists := v.index[relativePath]
	if !exists {
		return ErrNotExist
	}

	file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (v *VDrive) ApendFile(relativePath string, data []byte) error {
	relativePath = RelativePath(relativePath)

	fullpath, exists := v.index[relativePath]
	if !exists {
		return ErrNotExist
	}

	file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
