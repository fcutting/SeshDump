package processes

import (
    "fmt"
    "log"
    "encoding/xml"
    "encoding/json"
    "io/ioutil"
)

func outputScreen(processes []ProcessInfo) {
    for _, process := range processes {
        fmt.Println("PID       : ", process.PID)
        fmt.Println("Path      : ", process.Path)
        fmt.Println("Name      : ", process.Name)
        fmt.Println("User      : ", process.User)
        fmt.Println("SID       : ", process.SID)
        fmt.Println("Arguments : ", process.Arguments)
        fmt.Println("")
    }
}

func outputXML(processes []ProcessInfo, filename string) {
    file, err := xml.MarshalIndent(processes, "", "  ")
    
    if err != nil {
        log.Fatal("processes.outputXML() xml.MarshalIndent: ", err)
    }
    
    err = ioutil.WriteFile(filename, file, 0644)
    
    if err != nil {
        log.Fatal("processes.outputXML() ioutil.WriteFile: ", err)
    }
}

func outputJSON(processes []ProcessInfo, filename string) {
    file, err := json.MarshalIndent(processes, "", "  ")
    
    if err != nil {
        log.Fatal("processes.outputJSON() json.MarshalIndent: ", err)
    }
    
    err = ioutil.WriteFile(filename, file, 0644)
    
    if err != nil {
        log.Fatal("processes.outputJSON() ioutil.WriteFile: ", err)
    }
}
