package main

import (
    "./processes"
    "./winapi"
)

func main() {
    winapi.RtlAdjustPrivilege()
    processes.Dump()
    
    for {}
}
