package main

import (
    "os"
    "log"
    "time"
    
    "./arguments"
    "./processes"
    "./winapi"
)

func main() {
    // create session folder
    sessionFolder := "sessions/" + time.Now().Format("20060102150405") + "/"
    
    if _, err := os.Stat(sessionFolder); os.IsNotExist(err) {
        err := os.MkdirAll(sessionFolder, 0755)
        
        if err != nil {
            log.Fatal("main_os.MkdirAll: ", err)
        }
    }
    
    // parse arguments
    arguments := arguments.Parse(os.Args)
    
    // elevate privileges to SE_DEBUG_PRIVILEGE
    var previousValue bool
    
    err := winapi.RtlAdjustPrivilege(winapi.SE_DEBUG_PRIVILEGE, 1, 0, previousValue)
    
    if err != nil {
        log.Fatal("main_RtlAdjustPrivilege:", err)
    }
    
    // dump environment artifacts
    if arguments.Processes {
        processes.Dump(sessionFolder, arguments)
    }
}
