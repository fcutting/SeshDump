package winapi

import (
    "syscall"
    "unsafe"
)

var (
    dllPsapi                 = syscall.NewLazyDLL("psapi.dll")
    procEnumProcesses        = dllPsapi.NewProc("EnumProcesses")
    procGetModuleFileNameExA = dllPsapi.NewProc("GetModuleFileNameExA")
)

func EnumProcesses(buffer *uint32, bufferSize uint32, needed *uint32) error {
    r1, _, err := procEnumProcesses.Call(uintptr(unsafe.Pointer(buffer)), uintptr(bufferSize * 4), uintptr(unsafe.Pointer(needed)))
    
    if r1 == 0 {
        return err
    }
    
    return nil
}

func GetModuleFileNameExA(handle uintptr, buffer *byte, bufferSize uint32) (uintptr, error) {
    r1, _, err := procGetModuleFileNameExA.Call(handle, uintptr(0), uintptr(unsafe.Pointer(buffer)), uintptr(bufferSize))
    
    if r1 == 0 {
        return r1, err
    }
    
    return r1, nil
}
