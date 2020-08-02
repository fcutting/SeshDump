package winapi

import (
    "syscall"
    "unsafe"
    "log"
)

var (
    dllKernel32      = syscall.NewLazyDLL("kernel32.dll")
    procGetLastError = dllKernel32.NewProc("GetLastError")
    procOpenProcess  = dllKernel32.NewProc("OpenProcess")
    procCloseHandle  = dllKernel32.NewProc("CloseHandle")
    
    dllNtdll               = syscall.NewLazyDLL("ntdll.dll")
    procRtlAdjustPrivilege = dllNtdll.NewProc("RtlAdjustPrivilege")
    
    dllPsapi                 = syscall.NewLazyDLL("psapi.dll")
    procEnumProcesses        = dllPsapi.NewProc("EnumProcesses")
    procGetModuleFileNameExA = dllPsapi.NewProc("GetModuleFileNameExA")
)

func GetLastError(step string, err error) {
    r1, _, _ := procGetLastError.Call()
    
    if r1 > 0 {
        log.Fatal(step + ":", err)
    }
}

func RtlAdjustPrivilege() {
    var previousValue bool
    
    _, _, err := procRtlAdjustPrivilege.Call(uintptr(SE_DEBUG_PRIVILEGE), uintptr(1), uintptr(0), uintptr(unsafe.Pointer(&previousValue)))
    GetLastError("RtlAdjustPrivilege", err)
}

func EnumProcesses(buffer []uint32, bufferSize int, pidsWritten *uint32) {
    _, _, err := procEnumProcesses.Call(uintptr(unsafe.Pointer(&buffer[0])), uintptr(bufferSize * 4), uintptr(unsafe.Pointer(&pidsWritten)))
    GetLastError("EnumProcesses", err)
}

func OpenProcess(pid uint32) uintptr {
    handle, _, err := procOpenProcess.Call(uintptr(STANDARD_RIGHTS_REQUIRED) | uintptr(SYNCHRONIZE) | uintptr(0xFFFF), uintptr(1), uintptr(pid))
    GetLastError("OpenProcess", err)
    return handle
}

func CloseHandle(handle uintptr) {
    _, _, err := procCloseHandle.Call(handle)
    GetLastError("CloseHandle", err)
}

func GetModuleFileNameExA(handle uintptr, buffer []byte) uintptr {
    pathLength, _, err := procGetModuleFileNameExA.Call(handle, 0, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
    GetLastError("GetModuleFileNameExA", err)
    
    return pathLength
}
