package processes

import (
    "log"
    
    "../winapi"
    "../arguments"
)

type ProcessInfo struct {
    PID       int    `xml:"PID"`
    Path      string `xml:"Path"`
    Name      string `xml:"Name"`
    User      string `xml:"User"`
    SID       string `xml:"SID"`
    Arguments string `xml:"Arguments"`
}

func Dump(sessionFolder string, arguments arguments.Arguments) {
    // retrieve running processes information
    pids := getPIDs()
    var processInfo ProcessInfo
    processes := make([]ProcessInfo, 0)
    
    for _, pid := range pids {
        handle, err := winapi.OpenProcess(winapi.PROCESS_ALL_ACCESS, uint32(1), pid)
        
        if err != nil && err.Error() != "Access is denied." {
            log.Fatal("processes.Dump_OpenProcess: ", err)
        }
        
        if handle > 0 {
            processInfo.PID                   = int(pid)
            processInfo.Path                  = getPath(handle)
            processInfo.Name                  = getName(processInfo.Path)
            processInfo.User, processInfo.SID = getUser(handle)
            processInfo.Arguments             = getArguments(handle)
            
            processes = append(processes, processInfo)
            
            err = winapi.CloseHandle(handle)
            
            if err != nil {
                log.Fatal("processes.Dump_CloseHandle: ", err)
            }
        }
    }
    
    // export information
    if arguments.OutputScreen {
        outputScreen(processes)
    }
    
    if arguments.OutputXML {
        outputXML(processes, sessionFolder + "processes.xml")
    }
    
    if arguments.OutputJSON {
        outputJSON(processes, sessionFolder + "processes.json")
    }
}
