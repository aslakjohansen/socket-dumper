package main

import (
    "fmt"
    "os"
    "io"
    "time"
    "net"
    "bufio"
)

func append2log (log *os.File, t0 time.Time, t1 time.Time, line string) {
    var fullline string = fmt.Sprintf("%d.%09d,%d.%09d,%s\n", t0.Unix(), t0.Nanosecond(), t1.Unix(), t1.Nanosecond(), line)
    
    if _, err := log.Write([]byte(fullline)) ; err != nil {
        fmt.Println(err)
        os.Exit(5)
    }
    log.Sync()
}

func main () {
    // guard: command line args
    if len(os.Args) != 3 {
        fmt.Printf("Syntax: %s PORT FILENAME\n", os.Args[0])
        fmt.Printf("        %s 8000 mylog.csv\n", os.Args[0])
        os.Exit(1)
    }
    var port     string = os.Args[1]
    var filename string = os.Args[2]
    
    // open file
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(fmt.Sprintf("FATAL: Unable to open '%s' for appending: %s", filename, err))
        os.Exit(4)
    }
    defer f.Close()
    
    // listen
    listen, err := net.Listen("tcp4", ":"+port)
    if err != nil {
        fmt.Println(fmt.Sprintf("FATAL: Unable to listen to port %s: %s", port, err))
        os.Exit(2)
    }
    defer listen.Close()
    fmt.Println("Listening to port "+port)
    
    // service loop
    for {
        // accept
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println(fmt.Sprintf("FATAL: Unable to accept on port %s: %s", port, err))
            os.Exit(3)
        }
        
        // create reader
        reader := bufio.NewReader(conn)
        
        // session loop
        for {
            var t0 time.Time = time.Now()
            buffer, isPrefix, err := reader.ReadLine()
            var t1 time.Time = time.Now()
            
            // guard: err
            if err != nil {
                if err != io.EOF {
                    fmt.Println(fmt.Sprintf("ERROR: Unable to read connection on port %s: %s", port, err))
                }
                break
            }
            
            // isPrefix
            if isPrefix {
                fmt.Println(fmt.Sprintf("Warning: Truncating line received on port %s.", port))
            }
            
            // log to disk
            append2log(f, t0, t1, string(buffer))
        }
    }
}

