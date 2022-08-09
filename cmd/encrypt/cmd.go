package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"filippo.io/age"
)

const (
	publicKey = "age16r0huq2lu2u2x075yuhsj8fpm6z2ws2ndcxc75tzc85zg0uds5csmh5cq6"
	filePath  = "./files/a.txt"
	destPath  = "./encfiles/"
)

func main() {
	recipient, err := age.ParseX25519Recipient(publicKey)
	if err != nil {
		log.Println("Unable parse public key:", err.Error())
		return
	}

	srcFileIo, err := os.Open(filePath)
	if err != nil {
		log.Println("Unable open source file")
		return
	}

	destFile := filepath.Join(destPath, filepath.Base(filePath)+".age")
	destFileIo, err := os.Create(destFile)
	if err != nil {
		log.Println("Unable create dest file")
		return
	}

	wc, err := age.Encrypt(destFileIo, recipient)
	if err != nil {
		log.Println("Unable decrypt")
		return
	}

	_, err = io.Copy(wc, srcFileIo)
	if err != nil {
		log.Println("Unable Encrypt File")
		return
	}

	if err := wc.Close(); err != nil {
		log.Println("Unable Close Writer File")
		return
	}

	if err := destFileIo.Close(); err != nil {
		log.Println("Unable Close Encrypted File")
		return
	}

	if err := srcFileIo.Close(); err != nil {
		log.Println("Unable Close Source File")
		return
	}

	log.Println("File successfully encrypted")
}
