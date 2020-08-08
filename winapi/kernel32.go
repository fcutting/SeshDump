package winapi

import (
    "syscall"
)

var (
    dllKernel32     = syscall.NewLazyDLL("kernel32.dll")
    procOpenProcess = dllKernel32.NewProc("OpenProcess")
    procCloseHandle = dllKernel32.NewProc("CloseHandle")
)

func OpenProcess(desiredAccess uintptr, inheritHandle uint32, pid uint32) (uintptr, error) {
    r1, _, err := procOpenProcess.Call(desiredAccess, uintptr(inheritHandle), uintptr(pid))
    
    if r1 == 0 {
        return r1, err
    }
    
    return r1, nil
}

func CloseHandle(handle uintptr) error {
    r1, _, err := procCloseHandle.Call(handle)
    
    if r1 == 0 {
        return err
    }
    
    return nil
}
