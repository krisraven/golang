// to build an executable it should be called main
// this needs to be run in the GOROOT path (eg /usr/local/opt/go/libexec/src/shopping)
// echo $GOROOT

package main

import (
    "shopping"
    "fmt"
)

func main() {
    fmt.Println(shopping.PriceCheck(4343))
}