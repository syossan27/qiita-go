package main

import (
    "fmt"
    "os"

    "github.com/skratchdot/open-golang/open"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "qiita-go [word]")
        os.Exit(1)
    }
    word := os.Args[1]
    url := "http://qiita.com/search?utf8=✓&sort=rel&q=" + word
    open.Run(url)
}
