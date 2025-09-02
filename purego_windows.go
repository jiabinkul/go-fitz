//go:build (!cgo || nocgo) && windows

package fitz

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"
)

// EnsureDLLLoaded loads the DLL manually. Must be called before using other APIs.
func EnsureDLLLoaded(path string) error {
	if libmupdf != 0 {
		return nil // already loaded
	}

	handle, err := syscall.LoadLibrary(path)
	if err != nil {
		return fmt.Errorf("cannot load library %s: %w", path, err)
	}
	libmupdf = uintptr(handle)
	return nil
}

// procAddress returns the address of symbol name.
func procAddress(procName string) uintptr {
	if libmupdf == 0 {
		panic("libmupdf not loaded. Call EnsureDLLLoaded(path) first.")
	}
	addr, err := windows.GetProcAddress(windows.Handle(libmupdf), procName)
	if err != nil {
		panic(fmt.Errorf("cannot get proc address for %s: %w", procName, err))
	}
	return addr
}
