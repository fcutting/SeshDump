package winapi

import (
    "syscall"
    "unsafe"
)

var (
    dllNtdll                      = syscall.NewLazyDLL("ntdll.dll")
    procRtlAdjustPrivilege        = dllNtdll.NewProc("RtlAdjustPrivilege")
    procNtQueryInformationProcess = dllNtdll.NewProc("NtQueryInformationProcess")
    procNtReadVirtualMemory       = dllNtdll.NewProc("NtReadVirtualMemory")
)

func RtlAdjustPrivilege(privilege uint32, enablePrivilege uint32, isThreadPrivilege uint32, previousValue bool) error {
    r1, _, err := procRtlAdjustPrivilege.Call(uintptr(privilege), uintptr(enablePrivilege), uintptr(isThreadPrivilege), uintptr(unsafe.Pointer(&previousValue)))
    
    if r1 != 0 {
        return err
    }
    
    return nil
}

func NtQueryInformationProcess(handle uintptr, processInformationClass uintptr, processInformation *byte, processInformationSize uint32, needed *uint32) error {
    r1, _, err := procNtQueryInformationProcess.Call(handle, processInformationClass, uintptr(unsafe.Pointer(processInformation)), uintptr(processInformationSize), uintptr(unsafe.Pointer(needed)))
    
    if r1 != 0 {
        return err
    }
    
    return nil
}

func NtReadVirtualMemory(handle uintptr, baseAddress uintptr, buffer *byte, bufferSize uint32, needed *uint32) error {
    r1, _, err := procNtReadVirtualMemory.Call(handle, baseAddress, uintptr(unsafe.Pointer(buffer)), uintptr(bufferSize), uintptr(unsafe.Pointer(needed)))
    
    if r1 != 0 {
        return err
    }
    
    return nil
}
