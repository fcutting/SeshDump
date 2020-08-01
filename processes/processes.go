package processes

import (
    "fmt"
    "../winapi"
)

func Dump() {
    pids := getPIDs()
    
    fmt.Println(pids)
}

func getPIDs() []uint32 {
    bufferSize := 500
    buffer := make([]uint32, bufferSize)
    var pidsWritten *uint32
    
    // start at 500 and increase buffer by 50 until we get every PID
    for {
        winapi.EnumProcesses(buffer, bufferSize, pidsWritten)
        
        if buffer[len(buffer) - 1] == 0 {
            break
        }
        
        bufferSize += 50
        buffer = make([]uint32, bufferSize)
    }
    
    pidCount := 0
    
    for i := 1; i < bufferSize; i++ {
        if buffer[i] == 0 {
            pidCount = i
            break
        }
    }
    
    return buffer[:pidCount]
}
