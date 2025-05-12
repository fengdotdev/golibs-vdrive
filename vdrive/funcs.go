package vdrive

import "path"

func GenerateID() identifier {
	// Generate a unique ID for the VDrive instance
	// This is a placeholder implementation
	return "unique-id"
}

func RelativePath(rawpath string) string {
	relativePath := path.Clean(rawpath)
	return relativePath
}
