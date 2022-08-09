package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"filippo.io/age"
)

const (
	secretKey   = "AGE-SECRET-KEY-1RL468UQD20QVUTVPSUJLS3VPAVJJTF794K202CHDR4G4QLKNVJHQFCVZDP"
	encFile     = "./encfiles/a.txt.age"
	decryptPath = "./files/"
)

func main() {
	f, err := os.Open(encFile)
	if err != nil {
		log.Println("unable open file:", err.Error())
		return
	}

	fileName := filepath.Base(encFile)
	fileName = strings.Replace(fileName, ".age", "", -1)

	destDecFile := filepath.Join(decryptPath, fileName)
	fOut, err := os.Create(destDecFile)
	if err != nil {
		log.Println("Unable prepare destination file:", err.Error())
		return
	}

	identity, err := age.ParseX25519Identity(secretKey)
	if err != nil {
		log.Println("Unable parse identity:", err.Error())
		return
	}

	read, err := age.Decrypt(f, identity)
	if err != nil {
		log.Println("Error decrypt file:", err.Error())
		return
	}

	if _, err = io.Copy(fOut, read); err != nil {
		log.Println("Error process decrypt file", err.Error())
		return
	}

	log.Println("File successfully decrypted:", destDecFile)
}
