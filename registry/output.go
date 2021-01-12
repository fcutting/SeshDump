package registry

import (
	"fmt"
	"log"
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
)

func outputScreen(registries []RegistryInfo) {
	for _, registryInfo := range registries {
		fmt.Println("Path      : ", registryInfo.Path)
		fmt.Println("ValueName : ", registryInfo.ValueName)
		fmt.Println("Value     : ", registryInfo.Value)
		fmt.Println("")
	}
}

func outputXML(registries []RegistryInfo, filename string) {
	file, err := xml.MarshalIndent(registries, "", "  ")

	if err != nil {
		log.Fatal("registry.outputXML_MarshalIndent: ", err)
	}

	err = ioutil.WriteFile(filename, file, 0644)

	if err != nil {
		log.Fatal("registry.outputXML_ioutil.WriteFile: ", err)
	}
}

func outputJSON(registries []RegistryInfo, filename string) {
	file, err := json.MarshalIndent(registries, "", "  ")

	if err != nil {
		log.Fatal("registry.outputJSON_json.MarshalIndent: ", err)
	}

	err  = ioutil.WriteFile(filename, file, 0644)

	if err != nil {
		log.Fatal("registry.outputJSON_iotuil.WriteFile: ", err)
	}
}