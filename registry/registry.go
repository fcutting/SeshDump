package registry

import (
    "../arguments"
)

type RegistryInfo struct {
    Path      string `xml:"Path"`
    ValueName string `xml:"ValueName"`
    Value     string `xml:"Value"`
}

func Dump(sessionFolder string, arguments arguments.Arguments) {
    // retrieve registry information
    registries := dumpRun()
    registries = append(registries, dumpCLSIDReferences()...)
    registries = append(registries, dumpServices()...)

    // export information
    if arguments.OutputScreen {
        outputScreen(registries)
    }

    if arguments.OutputXML {
        outputXML(registries, sessionFolder + "registry.xml")
    }

    if arguments.OutputJSON {
        outputJSON(registries, sessionFolder + "registry.json")
    }
}

func find(slice []string, val string) (bool) {
    for _, item := range slice {
        if item == val {
            return true
        }
    }

    return false
}