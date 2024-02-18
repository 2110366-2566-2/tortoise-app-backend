//this is a simple program to hash a password using bcrypt

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    return string(bytes), err
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Password: ")
    password, _ := reader.ReadString('\n')

    hashedPassword, err := HashPassword(strings.TrimSpace(password))
    if err != nil {
        fmt.Println("Error hashing password:", err)
        return
    }

    fmt.Println("Hashed Password:", hashedPassword)
}