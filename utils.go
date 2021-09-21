package main

import (
	"path"
)

func replaceExtension(filePath string, newExt string) string {
	oldExt := path.Ext(filePath)
	return filePath[0:len(filePath)-len(oldExt)] + newExt
}
