package registry

import "golang.org/x/sys/windows/registry"

func dumpRun() []RegistryInfo {
    registries := make([]RegistryInfo, 0)
    
    // LOCAL_MACHINE keys
    keysLocalMachine := make([]string, 2)
    
    keysLocalMachine[0] = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
    keysLocalMachine[1] = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\RunOnce"
    
    // dump all values in each key
    for _, keyLocation := range keysLocalMachine {
        key, err := registry.OpenKey(registry.LOCAL_MACHINE, keyLocation, registry.QUERY_VALUE)
        
        if err != nil {
            continue
        }
        
        subvalueNames, err := key.ReadValueNames(0)
        
        if err != nil {
            continue
        }
        
        for _, valueName := range subvalueNames {
            value, _, err := key.GetStringValue(valueName)
            
            if err != nil {
                continue
            }
            
            var registryInfo RegistryInfo
            
            registryInfo.Path = "HKEY_LOCAL_MACHINE\\" + keyLocation
            registryInfo.ValueName = valueName
            registryInfo.Value = value
            
            registries = append(registries, registryInfo)
        }
    }
    
    // CURRENT_USER keys
    keysCurrentUser := make([]string, 2)
    
    keysCurrentUser[0] = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run"
    keysCurrentUser[1] = "SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\RunOnce"
    
    // dump all values in each key
    for _, keyLocation := range keysCurrentUser {
        key, err := registry.OpenKey(registry.CURRENT_USER, keyLocation, registry.QUERY_VALUE)
        
        if err != nil {
            continue
        }
        
        subvalueNames, err := key.ReadValueNames(0)
        
        if err != nil {
            continue
        }
        
        for _, valueName := range subvalueNames {
            value, _, err := key.GetStringValue(valueName)
            
            if err != nil {
                continue
            }
            
            var registryInfo RegistryInfo
            
            registryInfo.Path = "HKEY_CURRENT_USER\\" + keyLocation
            registryInfo.ValueName = valueName
            registryInfo.Value = value
            
            registries = append(registries, registryInfo)
        }
    }
    
    return registries
}

func dumpCLSIDReferences() []RegistryInfo {
    registries := make([]RegistryInfo, 0)
    
    key, err := registry.OpenKey(registry.CLASSES_ROOT, "CLSID", registry.ENUMERATE_SUB_KEYS)
    defer key.Close()
    
    if err != nil {
        return registries
    }
    
    subkeyNames, err := key.ReadSubKeyNames(0)
    
    if err != nil {
        return registries
    }
    
    for _, subkeyLocation := range subkeyNames {
        subkey, err := registry.OpenKey(registry.CLASSES_ROOT, "CLSID\\" + subkeyLocation, registry.ENUMERATE_SUB_KEYS)
        
        if err != nil {
            continue
        }
        
        subsubkeyNames, err := subkey.ReadSubKeyNames(0)
        
        if err != nil {
            continue
        }
        
        found := find(subsubkeyNames, "InprocServer32")
        
        if found {
            subkey, err := registry.OpenKey(registry.CLASSES_ROOT, "CLSID\\" + subkeyLocation + "\\InprocServer32", registry.QUERY_VALUE)
            
            if err != nil {
                continue
            }
            
            value, _, err := subkey.GetStringValue("")
            
            if err != nil {
                continue
            }
            
            var registryInfo RegistryInfo
            
            registryInfo.Path = "HKEY_CLASSES_ROOT\\CLSID\\" + subkeyLocation + "\\InprocServer32"
            registryInfo.ValueName = ""
            registryInfo.Value = value
            
            registries = append(registries, registryInfo)
        }
        
        found = find(subsubkeyNames, "LocalServer32")
        
        if found {
            subkey, err := registry.OpenKey(registry.CLASSES_ROOT, "CLSID\\" + subkeyLocation + "\\LocalServer32", registry.QUERY_VALUE)
            
            if err != nil {
                continue
            }
            
            value, _, err := subkey.GetStringValue("")
            
            if err != nil {
                continue
            }
            
            var registryInfo RegistryInfo
            
            registryInfo.Path = "HKEY_CLASSES_ROOT\\CLSID\\" + subkeyLocation + "\\LocalServer32"
            registryInfo.ValueName = ""
            registryInfo.Value = value
            
            registries = append(registries, registryInfo)
        }
    }
    
    return registries
}

func dumpServices() []RegistryInfo {
    registries := make([]RegistryInfo, 0)
    
    path := "SYSTEM\\CurrentControlSet\\Services"
    
    key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ENUMERATE_SUB_KEYS)
    defer key.Close()
    
    if err != nil {
        return registries
    }
    
    subkeyNames, err := key.ReadSubKeyNames(0)
    
    if err != nil {
        return registries
    }
    
    for _, subkeyLocation := range subkeyNames {
        // retrieve ImagePath value
        subkey, err := registry.OpenKey(registry.LOCAL_MACHINE, path + "\\" + subkeyLocation, registry.QUERY_VALUE)
    
        if err != nil {
            continue
        }
        
        value, _, err := subkey.GetStringValue("ImagePath")
        
        if err != nil {
            continue
        }
        
        var registryInfo RegistryInfo
        
        registryInfo.Path = path + "\\" + subkeyLocation
        registryInfo.ValueName = "ImagePath"
        registryInfo.Value = value
        
        registries = append(registries, registryInfo)
        
        // retrieve Parameters/ServiceDLL value
        subkey, err = registry.OpenKey(registry.LOCAL_MACHINE, path + "\\" + subkeyLocation, registry.ENUMERATE_SUB_KEYS)
    
        if err != nil {
            continue
        }
    
        subsubkeyNames, err := subkey.ReadSubKeyNames(0)
        
        if err != nil {
            continue
        }
        
        if find(subsubkeyNames, "Parameters") {
            subkey, err = registry.OpenKey(registry.LOCAL_MACHINE, path + "\\" + subkeyLocation + "\\Parameters", registry.QUERY_VALUE)
    
            value, _, err = subkey.GetStringValue("ServiceDLL")
    
            if err != nil {
                continue
            }
    
            var registryInfo RegistryInfo
    
            registryInfo.Path = path + "\\" + subkeyLocation + "\\Parameters"
            registryInfo.ValueName = "ServiceDLL"
            registryInfo.Value = value
    
            registries = append(registries, registryInfo)
        }
    }
    
    return registries
}