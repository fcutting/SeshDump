package processes

import (
    "fmt"
    "strings"
    // "unsafe"
    
    "../winapi"
    // "github.com/shirou/gopsutil/process"
)

type Process struct {
    PID       uint32
    Name      string
    Path      string
    User      string
    SID       string
}

func Dump() {
    pids := getPIDs()
    processes := make([]Process, len(pids))
    
    for i, pid := range pids {
        handle := winapi.OpenProcess(pid)
        
        if handle > 0 {
            defer winapi.CloseHandle(handle)
            
            processes[i].PID = pid
            processes[i].Path = getFilePath(handle)
            processes[i].Name = getFileName(processes[i].Path)
        }
    }
    
    for _, proc := range processes {
        if proc.PID > 0 {
            fmt.Println("pid:       ", proc.PID)
            fmt.Println("path:      ", proc.Path)
            fmt.Println("name:      ", proc.Name)
            fmt.Println("user:      ", proc.User)
            fmt.Println("sid:       ", proc.SID)
            fmt.Println("")
        }
    }
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

func getFilePath(handle uintptr) string {
    buffer := make([]byte, 260)
    
    pathLength := winapi.GetModuleFileNameExA(handle, buffer)
    
    return string(buffer[:int(pathLength)])
}

func getFileName(path string) string {
    name := strings.Split(path, "\\")
    return name[len(name) - 1]
}
