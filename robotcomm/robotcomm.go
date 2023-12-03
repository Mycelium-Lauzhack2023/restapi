// robot communication for the web server will be done as a tcp client
// would've been the dream to use this sadge
package robotcomm

import (
    "fmt"
    "net"
)

func ConnectToServer(address string, port string) {
    conn, err := net.Dial("tcp", address+":"+port)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Connected to server at", conn.RemoteAddr())
    // Handle the connection, e.g., read and write data
}

