package gfs_test

import (
	"runtime"
	"testing"

	"github.com/shockerli/gfs"
)

func TestCallerFunc(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"CallerFunc", "TestCallerFunc"},
	}
	for _, tt := range tests {
		if got := gfs.CallerFunc(); got != tt.want {
			t.Errorf("%q. CallerFunc() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCallerPkg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"CallerPkg", "command-line-arguments_test"},
	}
	for _, tt := range tests {
		if got := gfs.CallerPkg(); got != tt.want {
			t.Errorf("%q. CallerPkg() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCallerFile(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	tests := []struct {
		name string
		want string
	}{
		{"CallerFile", file},
	}
	for _, tt := range tests {
		if got := gfs.CallerFile(); got != tt.want {
			t.Errorf("%q. CallerFile() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
