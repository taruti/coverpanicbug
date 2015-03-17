package coverpanicbug

import "fmt"

func recoverAndPrint(f func()) {
        defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    f()
}

func CoverIsWeird() {
    a,b := 0,1
    panic("panic, does weird things to cover!")
    a += b
}


