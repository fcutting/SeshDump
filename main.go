package main

import (
    "fmt"
    "log"
    
    "./processes"
    "./winapi"
)

func main() {
    var previousValue bool
    
    err := winapi.RtlAdjustPrivilege(winapi.SE_DEBUG_PRIVILEGE, 1, 0, previousValue)
    
    if err != nil {
        log.Fatal("main_RtlAdjustPrivilege:", err)
    }
    
    processes := processes.Dump()
    
    for _, process := range processes {
        fmt.Println("PID      : ", process.PID)
        fmt.Println("Path     : ", process.Path)
        fmt.Println("Name     : ", process.Name)
        fmt.Println("User     : ", process.User)
        fmt.Println("SID      : ", process.SID)
        fmt.Println("Arguments: ", process.Arguments)
        fmt.Println("")
    }
}
