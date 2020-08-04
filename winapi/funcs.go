package winapi

import (
    "syscall"
    "golang.org/x/sys/windows"
    "unsafe"
    "log"
)

var (
    dllKernel32             = syscall.NewLazyDLL("kernel32.dll")
    procGetLastError        = dllKernel32.NewProc("GetLastError")
    procOpenProcess         = dllKernel32.NewProc("OpenProcess")
    procCloseHandle         = dllKernel32.NewProc("CloseHandle")
    
    dllNtdll                      = syscall.NewLazyDLL("ntdll.dll")
    procRtlAdjustPrivilege        = dllNtdll.NewProc("RtlAdjustPrivilege")
    
    dllPsapi                 = syscall.NewLazyDLL("psapi.dll")
    procEnumProcesses        = dllPsapi.NewProc("EnumProcesses")
    procGetModuleFileNameExA = dllPsapi.NewProc("GetModuleFileNameExA")
    
    dllAdvapi32                = syscall.NewLazyDLL("Advapi32.dll")
    procOpenProcessToken       = dllAdvapi32.NewProc("OpenProcessToken")
    procGetTokenInformation    = dllAdvapi32.NewProc("GetTokenInformation")
    procLookupAccountSidW      = dllAdvapi32.NewProc("LookupAccountSidW")
    procConvertSidToStringSidW = dllAdvapi32.NewProc("ConvertSidToStringSidW")
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

func OpenProcessToken(handle uintptr) uintptr {
    var thandle uintptr
    
    _, _, err := procOpenProcessToken.Call(handle, uintptr(TOKEN_QUERY), uintptr(unsafe.Pointer(&thandle)))
    GetLastError("OpenProcessToken", err)
    
    return thandle
}

func GetTokenInformation(thandle uintptr) *TOKEN_USER {
    n := uint32(50)
    b := make([]byte, n)
    
    _, _, _ = procGetTokenInformation.Call(thandle, uintptr(1), uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(unsafe.Pointer(&n)))
    
    r := unsafe.Pointer(&b[0])
    
    return (*TOKEN_USER)(r)
}

func ConvertSidToStringSidW(sid *SID) string {
    var s *uint16
    
    _, _, _ = procConvertSidToStringSidW.Call(uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(&s)))
    
    return windows.UTF16PtrToString(s)
}

func LookupAccountSidW(sid *SID) string {
    maxLength := uint32(256)
    var account *uint32
    
    name := make([]uint16, maxLength)
    domain := make([]uint16, maxLength)
    
    procLookupAccountSidW.Call(uintptr(0), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(&name[0])), uintptr(unsafe.Pointer(&maxLength)), uintptr(unsafe.Pointer(&domain[0])), uintptr(unsafe.Pointer(&maxLength)), uintptr(unsafe.Pointer(&account)))
    
    return syscall.UTF16ToString(domain) + "\\" + syscall.UTF16ToString(name)
}
