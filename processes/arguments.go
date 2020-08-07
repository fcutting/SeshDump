package processes

import (
    "log"
    "syscall"
    "unsafe"
    
    "../winapi"
)

func getArguments(handle uintptr) string {
    // get peb address
    var procBasicInfo winapi.PROCESS_BASIC_INFORMATION
    var needed uint32
    
    err := winapi.NtQueryInformationProcess(handle, uintptr(0), (*byte)(unsafe.Pointer(&procBasicInfo)), uint32(unsafe.Sizeof(procBasicInfo)), &needed)
    
    if err != nil {
        log.Fatal("processes.getArguments_NtQueryInformationProcess: ", err)
    }
    
    if needed == 0 {
        log.Fatal("processes.getArguments_NtQueryInformationProcess: Returned no data")
    }
    
    // get user proc params address
    var paramsAddress uintptr
    
    err = winapi.NtReadVirtualMemory(handle, procBasicInfo.PebBaseAddress + 32, (*byte)(unsafe.Pointer(&paramsAddress)), uint32(8), &needed)
    
    if err != nil {
        log.Fatal("processes.getArguments_NtReadVirtualMemory: ", err)
    }
    
    if needed == 0 {
        log.Fatal("processes.getArguments_NtReadVirtualMemory: Returned no data")
    }
    
    // get commandline unicode string structure
    var commandLine winapi.UNICODE_STRING
    
    err = winapi.NtReadVirtualMemory(handle, paramsAddress + 112, (*byte)(unsafe.Pointer(&commandLine)), uint32(16), &needed)
    
    if err != nil {
        log.Fatal("processes.getArguments_NtReadVirtualMemory: ", err)
    }
    
    if needed == 0 {
        log.Fatal("processes.getArguments_NtReadVirtualMemory: Returned no data")
    }
    
    // get commandline arguments
    cmd := make([]uint16, commandLine.Length)
    
    err = winapi.NtReadVirtualMemory(handle, commandLine.Buffer, (*byte)(unsafe.Pointer(&cmd[0])), uint32(commandLine.Length), &needed)
    
    if err != nil {
        log.Fatal("processes.getArguments_NtReadVirtualMemory: ", err)
    }
    
    if needed == 0 {
        log.Fatal("processes.getArguments_NtReadVirtualMemory: Returned no data")
    }
    
    return syscall.UTF16ToString(cmd)
}
