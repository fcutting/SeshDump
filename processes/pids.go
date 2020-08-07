package processes

import (
    "log"
    
    "../winapi"
)

func getPIDs() []uint32 {
    bufferSize := uint32(500)
    buffer     := make([]uint32, bufferSize)
    var needed uint32
    
    for {
        err := winapi.EnumProcesses(&buffer[0], bufferSize, &needed)
        
        if err != nil {
            log.Fatal("processes.getPIDs_EnumProcesses: ", err)
        }
        
        if buffer[len(buffer) - 1] == 0 {
            break
        }
        
        bufferSize += 50
        buffer = make([]uint32, bufferSize)
    }
    
    if needed == 0 {
        log.Fatal("processes.getPIDs_EnumProcesses: Returned no data")
    }
    
    var count int
    
    for i, pid := range buffer[1:] {
        if pid == 0 {
            count = i + 1
            break
        }
    }
    
    return buffer[1:count]
}
