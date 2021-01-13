package services

import (
    "golang.org/x/sys/windows/svc/mgr"
    "log"

    "../arguments"
)

type ServiceInfo struct {
    Name      string `xml:"Name"`
    Path      string `xml:"Path"`
    Type      string `xml:"Type"`
}

func Dump(sessionFolder string, arguments arguments.Arguments) {
    // retrieve services information
    services := make([]ServiceInfo, 0)
    
    mgrHandle, err := mgr.Connect()

    if err != nil {
        log.Fatal("services.Dump() mgr.Connect: ", err)
    }
    
    servicesNames, err := mgrHandle.ListServices()

    if err != nil {
        log.Fatal("services.Dump() mgrHandle.ListServices: ", err)
    }
    
    for _, serviceName := range servicesNames {
        service, err := mgrHandle.OpenService(serviceName)
        
        if err != nil {
            continue
        }
        
        config, err := service.Config()
    
        if err != nil {
            continue
        }
        
        var serviceInfo ServiceInfo
        
        serviceInfo.Name = config.DisplayName
        serviceInfo.Path = config.BinaryPathName
        serviceInfo.Type = serviceTypeToString(config.ServiceType)
        
        services = append(services, serviceInfo)
    }
    
    // export information
    if arguments.OutputScreen {
        outputScreen(services)
    }
    
    if arguments.OutputXML {
        outputXML(services, sessionFolder + "services.xml")
    }
    
    if arguments.OutputJSON {
        outputJSON(services, sessionFolder + "services.json")
    }
}

func serviceTypeToString(serviceType uint32) string {
    serviceTypeString := ""
    
    switch (serviceType) {
    case 1:
        serviceTypeString = "KernelDriver"
    case 2:
        serviceTypeString = "FileSystemDriver"
    case 4:
        serviceTypeString = "Adapter"
    case 8:
        serviceTypeString = "RecognizerDriver"
    case 16:
        serviceTypeString = "Win32OwnProcess"
    case 32:
        serviceTypeString = "Win32ShareProcess"
    case 256:
        serviceTypeString = "InteractiveProcess"
    }
    
    return serviceTypeString
}