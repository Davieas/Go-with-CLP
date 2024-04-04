package communication

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func TcpConnection() {
    // Prompt the user to enter the IP address of the logical programmer controller
    fmt.Print("Enter the IP address of the logical programmer controller: ")
    reader := bufio.NewReader(os.Stdin)
    lpcIP, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        os.Exit(1)
    }
    lpcIP = strings.TrimSpace(lpcIP)

    // Define the port of the logical programmer controller
    lpcPort := "8080"

    // Construct the address of the logical programmer controller
    lpcAddr := fmt.Sprintf("%s:%s", lpcIP, lpcPort)

    // Connect to the logical programmer controller
    conn, err := net.Dial("tcp", lpcAddr)
    if err != nil {
        fmt.Println("Error connecting:", err)
        os.Exit(1)
    }
    defer conn.Close()

    fmt.Println("Connected to Logical Programmer Controller")

  
}
