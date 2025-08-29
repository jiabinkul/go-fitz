//go:build (!cgo || nocgo) && windows

package fitz

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"
)

// 自定义 DLL 目录
var customDLLDir string

// SetDLLDir 设置 libmupdf.dll 搜索目录
func SetDLLDir(dir string) {
	customDLLDir = dir
}

// loadLibrary loads the dll and panics on error.
func loadLibrary() uintptr {
	dllPath := "libmupdf.dll"

	// 如果指定了自定义目录，则使用完整路径
	if customDLLDir != "" {
		dllPath = filepath.Join(customDLLDir, "libmupdf.dll")
	}

	// 检查文件是否存在
	if _, err := os.Stat(dllPath); os.IsNotExist(err) {
		panic(fmt.Errorf("libmupdf.dll 不存在: %s", dllPath))
	}

	handle, err := syscall.LoadLibrary(dllPath)
	if err != nil {
		panic(fmt.Errorf("无法加载库 %s: %w", dllPath, err))
	}

	return uintptr(handle)
}

// procAddress returns the address of symbol name.
func procAddress(handle uintptr, procName string) uintptr {
	addr, err := windows.GetProcAddress(windows.Handle(handle), procName)
	if err != nil {
		panic(fmt.Errorf("cannot get proc address for %s: %w", procName, err))
	}

	return addr
}
