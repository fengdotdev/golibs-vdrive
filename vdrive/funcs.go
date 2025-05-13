package vdrive

import (
	"path"

	"github.com/fengdotdev/golibs-helperfuncs/datum"

	"github.com/fengdotdev/golibs-helperfuncs/unique"
)

func GenerateID() identifier {
	id := unique.RamdomUUID()
	return id
}

func HashStream(stream string) string {
	panic("not implemented")
	return "hashed-stream"
}

func HashFile(filePath string) (string, error) {
	panic("not implemented")
	return "hashed-file", nil
}

func HashData(data []byte) string {
	result := datum.GetSHA256Bytes(data)
	return result
}

func RelativePath(rawpath string) string {
	relativePath := path.Clean(rawpath)
	return relativePath
}
