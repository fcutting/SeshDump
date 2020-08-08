package output

import (
    "fmt"
    "log"
    "encoding/xml"
    "encoding/json"
    "io/ioutil"
    
    "../processes"
)

func ProcessesScreen(processes []processes.ProcessInfo) {
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

func ProcessesXML(processes []processes.ProcessInfo) {
    file, err := xml.MarshalIndent(processes, "", "    ")
    
    if err != nil {
        log.Fatal("output.ProcessesXML_xml.MarshalIndent: ", err)
    }
    
    err = ioutil.WriteFile("processes.xml", file, 0644)
    
    if err != nil {
        log.Fatal("output.ProcessesXML_ioutil.WriteFile: ", err)
    }
}

func ProcessesJSON(processes []processes.ProcessInfo) {
    file, err := json.MarshalIndent(processes, "", "")
    
    if err != nil {
        log.Fatal("output.ProcessesJSON_json.MarshalIndent: ", err)
    }
    
    err = ioutil.WriteFile("processes.json", file, 0644)
    
    if err != nil {
        log.Fatal("output.ProcessesJSON_ioutil.WriteFile: ", err)
    }
}
