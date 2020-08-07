package winapi

import (
    "syscall"
    "unsafe"
    
    "golang.org/x/sys/windows"
)

var (
    dllAdvapi32                = syscall.NewLazyDLL("advapi32.dll")
    procOpenProcessToken       = dllAdvapi32.NewProc("OpenProcessToken")
    procGetTokenInformation    = dllAdvapi32.NewProc("GetTokenInformation")
    procConvertSidToStringSidW = dllAdvapi32.NewProc("ConvertSidToStringSidW")
    procLookupAccountSidW      = dllAdvapi32.NewProc("LookupAccountSidW")
)

func OpenProcessToken(handle uintptr, desiredAccess uintptr) (uintptr, error) {
    var thandle uintptr
    
    r1, _, err := procOpenProcessToken.Call(handle, desiredAccess, uintptr(unsafe.Pointer(&thandle)))
    
    if r1 == 0 {
        return thandle, err
    }
    
    return thandle, nil
}

func GetTokenInformation(thandle uintptr, tokenInformationClass uintptr) (*TOKEN_USER, error) {
    buffer := make([]byte, 50)
    var needed uint32
    
    r1, _, err := procGetTokenInformation.Call(thandle, tokenInformationClass, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)), uintptr(unsafe.Pointer(&needed)))
    
    if r1 == 0 {
        return nil, err
    }
    
    return (*TOKEN_USER)(unsafe.Pointer(&buffer[0])), nil
}

func ConvertSidToStringSidW(sid *SID) (string, error) {
    var stringSid *uint16
    
    r1, _, err := procConvertSidToStringSidW.Call(uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(&stringSid)))
    
    if r1 == 0 {
        return "", err
    }
    
    return windows.UTF16PtrToString(stringSid), nil
}

func LookupAccountSidW(systemName string, sid *SID) (string, error) {
    maxLength := uint32(256)
    var account *uint32
    
    name := make([]uint16, maxLength)
    domain := make([]uint16, maxLength)
    
    r1, _, err := procLookupAccountSidW.Call(uintptr(unsafe.Pointer(&systemName)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(&name[0])), uintptr(unsafe.Pointer(&maxLength)), uintptr(unsafe.Pointer(&domain[0])), uintptr(unsafe.Pointer(&maxLength)), uintptr(unsafe.Pointer(&account)))
    
    if r1 == 0 {
        return "", err
    }
    
    return syscall.UTF16ToString(domain) + "\\" + syscall.UTF16ToString(name), nil
}
