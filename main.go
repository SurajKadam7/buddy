package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var meta Meta

const defaultFilePrefix = "asmt"
const defaultFolderPrefix = "dir"
const defaultLocation = "/Users/suraj.kadam/go/src/bitbucket.org/junglee_games/assignments"

type Meta struct {
	Folder int `json:"folder,omitempty"`
	File   int `json:"file,omitempty"`
}

func customeTime() string {
	t := time.Now()
	y := t.Year() % ((t.Year() / 1000) * 1000)
	m := int(t.Month())
	d := t.Day()

	yy := padNumber(2, y)
	mm := padNumber(2, m)
	dd := padNumber(2, d)
	return fmt.Sprintf("%s.%s.%s", yy, mm, dd)
}

func padNumber(p, num int) string {
	format := fmt.Sprintf("%%0%dd", p)
	return fmt.Sprintf(format, num)
}

func customeNameWithSuffix(num int) string {
	t := customeTime()
	pn := padNumber(3, num)
	return fmt.Sprintf("%s_%s", t, pn)
}

func nameWithSuffixAndExtension(num int, extension string) string {
	cn := customeNameWithSuffix(num)
	return fmt.Sprintf("%s.%s", cn, extension)
}

func getDefault() (folder string, file string, err error) {
	metaPath := filepath.Join(defaultLocation, "meta")
	f, err := os.OpenFile(metaPath, os.O_RDWR, 0666)
	if os.IsNotExist(err) {
		f, err = os.Create(metaPath)
	}
	if err != nil {
		return "", "", err // Return error from function
	}

	err = json.NewDecoder(f).Decode(&meta)
	if err != nil && err != io.EOF { // Simplified EOF check
		return "", "", err
	}

	meta.File++
	meta.Folder = meta.File / 50
	f.Seek(0, 0)
	err = json.NewEncoder(f).Encode(&meta) // Encode updated meta
	if err != nil {
		return "", "", err
	}

	fileSuff := nameWithSuffixAndExtension(meta.File, "go")
	folderSuff := customeNameWithSuffix(meta.Folder)

	file = fmt.Sprintf("%s.%s", defaultFilePrefix, fileSuff)
	folder = fmt.Sprintf("%s.%s", defaultFolderPrefix, folderSuff)
	return
}

func main() {
	createDirIfNotExist(defaultLocation)
	folderName, fileName, err := getDefault()
	if err != nil {
		fmt.Println("Error getting defaults:", err) // Handle error from getDefault
		return
	}

	folderPath := filepath.Join(defaultLocation, folderName)
	createDirIfNotExist(folderPath)
	path := filepath.Join(folderPath, fileName)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err) // Handle create error
		return
	}
	funcData := fmt.Sprintf("package main; func %s(){}", getMainFunc())
	data := []byte(funcData)

	data, _ = format.Source(data)
	file.Write(data)
	file.Close() // Close the file even if Write doesn't error

	createMain(folderPath)
	cmd := exec.Command("code", folderPath)
	cmd.Run()
	cmd = exec.Command("code", path)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		log.Printf("not able to open code in vs code\nPlase use below the file path to write the code\npath : %s ", path)
	}
}

func getMainFunc() string {
	funcName := padNumber(3, meta.File)
	return fmt.Sprintf("main_%s", funcName)
}

func createMain(dirPath string) {
	path := filepath.Join(dirPath, "main.go")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	mainFunc := fmt.Sprintf("package main; func main(){%s()}", getMainFunc())
	data := []byte(mainFunc)
	data, _ = format.Source(data)
	f.Write(data)
	f.Close()
}

func createDirIfNotExist(path string) {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		return
	}

	if err := os.MkdirAll(path, 0775); err != nil {
		panic(err)
	}
}
