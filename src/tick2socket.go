package main

import (
    "fmt"
    "os"
    "net"
)

const (
    eol = "\n"
)

func main () {
    // guard: command line args
    if len(os.Args) != 2 {
        fmt.Printf("Syntax: %s PORT\n", os.Args[0])
        fmt.Printf("        %s 8000\n", os.Args[0])
        os.Exit(1)
    }
    var port     string = os.Args[1]
    
    // connect
    conn, err := net.Dial("tcp", ":"+port)
    if err != nil {
        fmt.Println(fmt.Sprintf("FATAL: Unable to connect to port %s: %s", port, err))
        os.Exit(2)
    }
    defer conn.Close()
    fmt.Println("Connected to port "+port)
    
    // service loop
    var counter int = 0
    for {
        // write
        conn.Write([]byte(fmt.Sprintf("%d", counter)))
        conn.Write([]byte(eol))
        
        fmt.Println("tick", counter)
        counter++
    }
}

