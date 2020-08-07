package main

import (
    "os"
    "log"
    
    "./arguments"
    "./processes"
    "./winapi"
    "./output"
)

func main() {
    arguments := arguments.Parse(os.Args)
    
    var previousValue bool
    
    err := winapi.RtlAdjustPrivilege(winapi.SE_DEBUG_PRIVILEGE, 1, 0, previousValue)
    
    if err != nil {
        log.Fatal("main_RtlAdjustPrivilege:", err)
    }
    
    if arguments.Processes {
        processes := processes.Dump()
        
        if arguments.OutputScreen {
            output.ProcessesScreen(processes)
        } else if arguments.OutputXML {
            output.ProcessesXML(processes)
        } else if arguments.OutputJSON {
            output.ProcessesJSON(processes)
        }
    }
}
