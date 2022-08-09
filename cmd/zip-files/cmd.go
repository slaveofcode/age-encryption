package main

import (
	"io"
	"log"
	"os"

	"github.com/yeka/zip"
)

type fileSrc struct {
	name string
	path string
}
type fileIo struct {
	name   string
	reader io.Reader
}

var files []fileSrc = []fileSrc{
	{"a.txt", "./files/a.txt"},
	{"b.txt", "./files/b.txt"},
	{"cool-image.png", "./files/img.png"},
}

func readFiles(fileSources []fileSrc) []fileIo {
	var ios []fileIo = []fileIo{}

	for _, item := range fileSources {
		f, _ := os.Open(item.path)
		ios = append(ios, fileIo{item.name, f})
	}

	return ios
}

func main() {
	password := "foobar"
	destZip := "./myfile.zip"

	fileIos := readFiles(files)
	fzip, err := os.Create(destZip)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()

	for _, fileIo := range fileIos {

		w, err := zipw.Encrypt(fileIo.name, password, zip.AES256Encryption)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(w, fileIo.reader)
		if err != nil {
			log.Fatal(err)
		}

		fileIo.reader.(*os.File).Close()
	}
	zipw.Flush()
}
