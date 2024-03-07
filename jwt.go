package main

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
)

func generateJWTSecret() (string, error) {
    secretLength := 64
    secretBytes := make([]byte, secretLength)
    _, err := rand.Read(secretBytes)
    if err != nil {
        return "", err
    }
    secret := base64.URLEncoding.EncodeToString(secretBytes)
    return secret, nil
}

func main() {
    secret, err := generateJWTSecret()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Generated HS512 JWT Secret:", secret)
}
