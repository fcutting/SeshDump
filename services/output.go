package services

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "log"
)

func outputScreen(services []ServiceInfo) {
    for _, serviceInfo := range services {
        fmt.Println("Name : ", serviceInfo.Name)
        fmt.Println("Path : ", serviceInfo.Path)
        fmt.Println("Type : ", serviceInfo.Type)
        fmt.Println("")
    }
}

func outputXML(services []ServiceInfo, filename string) {
    file, err := xml.MarshalIndent(services, "", "  ")
    
    if err != nil {
        log.Fatal("services.outputXML() xml.MarshalIndent: ", err)
    }
    
    err = ioutil.WriteFile(filename, file, 0644)
    
    if err != nil {
        log.Fatal("services.outputXML() ioutil.WriteFile: ", err)
    }
}

func outputJSON(services []ServiceInfo, filename string) {
    file, err := json.MarshalIndent(services, "", "  ")
    
    if err != nil {
        log.Fatal("services.outputJSON() json.MarshalIndent: ", err)
    }
    
    err  = ioutil.WriteFile(filename, file, 0644)
    
    if err != nil {
        log.Fatal("services.outputJSON() iotuil.WriteFile: ", err)
    }
}