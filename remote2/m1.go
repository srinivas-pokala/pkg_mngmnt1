package remote2

import "fmt"

func Print(str string) string {
        fmt.Printf("string str: %s\n", str)
        return str + "**"
}
