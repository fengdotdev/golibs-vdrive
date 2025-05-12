package vdrive

import "errors"

var (
	ErrNotExist   = errors.New("file does not exist")
	ErrPermission = errors.New("permission denied")
	ErrInvalid    = errors.New("invalid argument")
	ErrFileExists = errors.New("file already exists")

)
