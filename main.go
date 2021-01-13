package main

import (
    "os"
    "log"
    "time"
    
    "./arguments"
    "./processes"
    "./winapi"
    "./registry"
    "./services"
)

func main() {
    // parse arguments
    args := arguments.Parse(os.Args)

    // create session folder
    var sessionFolder string

    if args.OutputJSON || args.OutputXML {
        sessionFolder = "sessions/" + time.Now().Format("20060102150405") + "/"

        if _, err := os.Stat(sessionFolder); os.IsNotExist(err) {
            err := os.MkdirAll(sessionFolder, 0755)

            if err != nil {
                log.Fatal("main() os.MkdirAll: ", err)
            }
        }
    }
    
    // elevate privileges to SE_DEBUG_PRIVILEGE
    var previousValue bool
    
    err := winapi.RtlAdjustPrivilege(winapi.SE_DEBUG_PRIVILEGE, 1, 0, previousValue)
    
    if err != nil {
        log.Fatal("main() winapi.RtlAdjustPrivilege:", err)
    }
    
    // dump environment artifacts
    if args.Processes {
        processes.Dump(sessionFolder, args)
    }
    
    if args.Registry {
        registry.Dump(sessionFolder, args)
    }
    
    if args.Services {
        services.Dump(sessionFolder, args)
    }
}
