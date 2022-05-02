package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("asdasd")
//     file, err := os.Open("file")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()  // will give no such file or directory error. If this is commented out
                        // it will give file declared but not used error message
}