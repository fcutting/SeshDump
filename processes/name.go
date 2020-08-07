package processes

import (
    "strings"
)

func getName(path string) string {
    name := strings.Split(path, "\\")
    return name[len(name) - 1]
}
