package gfs

import (
	"path/filepath"
	"runtime"
	"strings"
)

// CallerFunc return the function of the caller
func CallerFunc() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name()
	nameEnd := filepath.Ext(nameFull)
	name := strings.TrimPrefix(nameEnd, ".")

	return name
}

// CallerPkg return the package of the caller
func CallerPkg() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name()
	sp := strings.Split(nameFull, ".")

	return sp[0]
}

// CallerFile return the file of the caller
func CallerFile() string {
	_, file, _, _ := runtime.Caller(1)

	return file
}
