package winapi

const(
    SE_DEBUG_PRIVILEGE = uint32(20)
    
    STANDARD_RIGHTS_REQUIRED = uintptr(0xF0000)
    SYNCHRONIZE              = uintptr(0x00100000)
    PROCESS_ALL_ACCESS       = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | uintptr(0xFFFF)
    
    TOKEN_QUERY    = uintptr(8)
    TIC_TOKEN_USER = uintptr(1)
)

type SID struct {}

type SID_AND_ATTRIBUTES struct {
    Sid        *SID
    Attributes uint32
}

type TOKEN_USER struct {
    User SID_AND_ATTRIBUTES
}

type PROCESS_BASIC_INFORMATION struct {
    Reserved1       uintptr
    PebBaseAddress  uintptr
    Reserved2       [2]uintptr
    UniqueProcessId uintptr
    Reserved3       uintptr
}

type UNICODE_STRING struct {
    Length    uint16
    MaxLength uint16
    Buffer    uintptr
}
