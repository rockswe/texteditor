package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("Simple Text Editor in Go")

    //wherever
    
    filename := "example.txt" // This can be taken from command-line arguments

    // Reading from a file
    content, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("Current Content:", string(content))

    // Simple text input
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter text to append:")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")

    // Appending to the file
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer f.Close()

    if _, err = f.WriteString(input + "\n"); err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("Updated file successfully.")
}
