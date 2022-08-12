package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Println("Welcome to Zippi")
	fmt.Println("---------------------")
	ZipWriter()
	fmt.Println("---------------------")
	fmt.Println("Done")
}

func ZipWriter() {

	baseFolder := "C:/ZippiTest/SomeFiles/"

	outFileWriter, err := os.Create(`C:/ZippiTest/zip.zip`)
	if err != nil {
		fmt.Println(err)
	}
	defer outFileWriter.Close()

	w := zip.NewWriter(outFileWriter)

	addFiles(w, baseFolder, "")

	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
