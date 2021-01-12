package arguments

type Arguments struct {
    Processes    bool
    Registry     bool
    Services     bool
    OutputScreen bool
    OutputXML    bool
    OutputJSON   bool
}

func Parse(args []string) Arguments {
    arguments := newArguments()
    
    for _, arg := range args {
        switch arg {
        case "-p":
            arguments.Processes = true
        case "-r":
            arguments.Registry = true
        case "-s":
            arguments.Services = true
        case "-oS":
            arguments.OutputScreen = true
        case "-oX":
            arguments.OutputXML = true
        case "-oJ":
            arguments.OutputJSON = true
        }
    }
    
    return arguments
}

func newArguments() Arguments {
    var arguments Arguments
    
    arguments.Processes    = false
    arguments.Registry     = false
    arguments.OutputScreen = false
    arguments.OutputXML    = false
    arguments.OutputJSON   = false
    
    return arguments
}
