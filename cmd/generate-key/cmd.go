package main

import (
  "fmt"
  "os"
  "log"
  "filippo.io/age"
)

func main() {
  identity, err := age.GenerateX25519Identity()
  if err != nil {
	log.Fatalf("Failed to generate key pair: %v", err)
  }

  log.Printf("Public key: %s\n", identity.Recipient().String())
  log.Printf("Private key: %s\n", identity.String())

  f, _ := os.Create("./key.txt")

  f.Write([]byte(fmt.Sprintf(`Public key: %s
Private key: %s`, identity.Recipient().String(), identity.String())))
  f.Close()

  log.Println("Key Created")
}
