package processes

import (
    "log"
    
    "../winapi"
)

func getPath(handle uintptr) string {
    buffer := make([]byte, 260)
    
    length, err := winapi.GetModuleFileNameExA(handle, &buffer[0], uint32(len(buffer)))
    
    if err != nil {
        log.Fatal("processes.getPath() winapi.GetModuleFileNameExA: ", err)
    }
    
    return string(buffer[:int(length)])
}
