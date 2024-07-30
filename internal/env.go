// Package internal /*
/*
//Copyright © 2024 github.com/1in9 HERE <892095205@qq.com>
//*/
package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

// Elevate
//
// Request administrator privileges to restart the program.
func Elevate() error {
	verb := "runas"
	file := os.Args[0]
	params := strings.Join(os.Args[1:], " ")
	dir := ""

	return ShellExecute(verb, file, params, dir, syscall.SW_HIDE)
}

// ShellExecute
//
// Use ShellExecuteW to elevate permissions.
func ShellExecute(verb, file, params, directory string, showCmd int) error {
	var (
		err error
	)

	hWnd := 0
	lpVerb, _ := syscall.UTF16PtrFromString(verb)
	lpFile, _ := syscall.UTF16PtrFromString(file)
	lpParameters, _ := syscall.UTF16PtrFromString(params)
	lpDirectory, _ := syscall.UTF16PtrFromString(directory)
	nShowCmd := uintptr(showCmd)

	shell32, err := syscall.LoadLibrary("shell32.dll")
	if err != nil {
		return fmt.Errorf("failed to load shell32.dll: %v", err)
	}
	defer func(handle syscall.Handle) {
		err := syscall.FreeLibrary(handle)
		if err != nil {

		}
	}(shell32)

	shellExecuteW, err := syscall.GetProcAddress(shell32, "ShellExecuteW")
	if err != nil {
		return fmt.Errorf("failed to get ShellExecuteW address: %v", err)
	}

	ret, _, _ := syscall.SyscallN(
		shellExecuteW,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpVerb)),
		uintptr(unsafe.Pointer(lpFile)),
		uintptr(unsafe.Pointer(lpParameters)),
		uintptr(unsafe.Pointer(lpDirectory)),
		nShowCmd,
	)

	if ret <= 32 {
		err = fmt.Errorf("ShellExecuteW failed with error code %d", ret)
	}

	return err
}

// IsElevated
//
// Check if it is running with administrator privileges.
func IsElevated() bool {
	_, err := os.Open(`\\.\PHYSICALDRIVE0`)
	return err == nil
}

// SetEnv
//
// Set system environment variables
func SetEnv(key, value string) error {
	if !IsElevated() {
		if err := Elevate(); err != nil {
			return err
		}
		os.Exit(0)
	}

	cmd := exec.Command("setx", "/M", key, value)
	return cmd.Run()
}

// AppendToPath 将路径添加到系统 Path 环境变量中（如果不存在）
func AppendToPath(path string) error {
	if !IsElevated() {
		if err := Elevate(); err != nil {
			return err
		}
		os.Exit(0)
	}

	currentPath, err := GetCurrentPath()
	if err != nil {
		return err
	}

	// 检查路径是否已存在于 Path 环境变量中
	if ContainsPath(currentPath, path) {
		return nil // 路径已存在，无需添加
	}

	newPath := appendPath(currentPath, path)

	// 执行 setx 命令
	cmd := exec.Command("setx", "/M", "Path", newPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing setx command: %v, output: %s", err, out)
	}

	return nil
}

// ContainsPath 检查路径是否已经存在于 Path 环境变量中
func ContainsPath(currentPath, path string) bool {
	for _, p := range strings.Split(currentPath, ";") {
		if strings.TrimSpace(p) == strings.TrimSpace(path) {
			return true
		}
	}
	return false
}

// GetCurrentPath 获取当前系统 Path 环境变量
func GetCurrentPath() (string, error) {
	cmd := exec.Command("set", "Path")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	paths := strings.Split(string(output), "\r\n")
	var currentPath string
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if strings.HasPrefix(p, "Path=") {
			currentPath = p[len("Path="):]
			break
		}
	}

	return currentPath, nil
}

// appendPath 将新路径添加到当前路径字符串中
func appendPath(currentPath, newPath string) string {
	if currentPath == "" {
		return newPath
	}

	// 如果当前路径以分号结束，则不需要额外的分号
	if !strings.HasSuffix(currentPath, ";") {
		currentPath += ";"
	}

	return currentPath + newPath
}

// RemovePath 从当前路径字符串中移除指定的路径
func RemovePath(currentPath, path string) error {
	if !IsElevated() {
		if err := Elevate(); err != nil {
			log.Fatalf("failed to elevate: %v", err)
		}
		os.Exit(0)
	}

	// 检查路径是否已存在于 Path 环境变量中
	if !ContainsPath(currentPath, path) {
		log.Fatalf("path %s not found in system variable path", path)
	}
	parts := strings.Split(currentPath, ";")
	var newPath string
	for _, p := range parts {
		if strings.TrimSpace(p) != strings.TrimSpace(path) {
			if newPath != "" {
				newPath += ";"
			}
			newPath += p
		}
	}
	cmd := exec.Command("setx", "/M", "Path", newPath)

	return cmd.Run()
}
