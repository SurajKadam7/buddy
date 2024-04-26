package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"time"
// )

// const customeMeta string = "custom_meta"
// const defaultMeta string = "default_meta"

// type Statergy int

// const defaultFilePrefix = "sample"
// const defaultFolderPrefix = "collector"
// const defaultLocation = "/Users/suraj.kadam/go/src/bitbucket.org/junglee_games"

// const (
// 	MultiFile Statergy = iota
// 	MultiFolder
// )

// type statergy func(meta meta) (path string, err error)

// func getStatergy(s Statergy) statergy {
// 	if s == MultiFile {
// 		return multiFile
// 	}
// 	return multiFolder
// }

// func customeTime() string {
// 	t := time.Now()
// 	y := t.Year() % ((t.Year() / 1000) * 1000)
// 	m := int(t.Month())
// 	d := t.Day()

// 	yy := padNumber(2, y)
// 	mm := padNumber(2, m)
// 	dd := padNumber(2, d)
// 	return fmt.Sprintf("%d.%d.%d", yy, mm, dd)
// }

// func padNumber(p, num int) string {
// 	format := fmt.Sprintf("%%0%dd", p)
// 	return fmt.Sprintf(format, num)
// }

// func customeNameWithSuffix(num int) string {
// 	t := customeTime()
// 	pn := padNumber(3, num)
// 	return fmt.Sprintf("%s_%s", t, pn)
// }

// func nameWithSuffixAndExtension(num int, extension string) string {
// 	cn := customeNameWithSuffix(num)
// 	return fmt.Sprintf("%s.%s", cn, extension)
// }

// func multiFile(meta meta) (path string, err error) {
// 	path = nameWithSuffixAndExtension(meta.FileNumber, "go")
// 	// opening file if exist if not then create and then open it
// 	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		return
// 	}
// 	defer f.Close()

// 	data, err := os.ReadFile("start.templ")
// 	if err != nil {
// 		return
// 	}

// 	_, err = f.Write(data)
// 	return
// }

// func multiFolder(meta meta) (path string, err error) {
// 	folderPath := customeNameWithSuffix(meta.FolderNumber)

// 	err = os.Mkdir(folderPath, 0775)
// 	if err != nil {
// 		return
// 	}

// 	path = filepath.Join(folderPath, "main.go")
// 	f, err := os.Create(path)
// 	if err != nil {
// 		return
// 	}
// 	defer f.Close()

// 	data, err := os.ReadFile("main.templ")
// 	if err != nil {
// 		return
// 	}

// 	_, err = f.Write(data)
// 	return
// }

// func isExist(path string) bool {
// 	_, err := os.Stat(path)
// 	// os.IsNotExist is used to check for the specific error of a non-existent file or directory
// 	return err == nil || os.IsExist(err)
// }

// func initIfNot(m meta) {
// 	path := filepath.Join(m.Location, "root", m.FolderName)
// 	if isExist(path) {
// 		return
// 	}

// 	// initialization flow
// 	err := os.Mkdir(path, 0775)
// 	if err != nil {
// 		panic(err)
// 	}

// 	metaPath := filepath.Join(path, "meta")
// 	err = creatFileIfNotExist(metaPath)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func creatFileIfNotExist(path string) (err error) {
// 	if isExist(path) {
// 		return
// 	}

// 	f, err := os.Create(path)
// 	if err != nil {
// 		return
// 	}

// 	m := &meta{
// 		FileName:   "asmt",
// 		FolderName: "dir",
// 		Location: ,
// 	}

// 	f.Close()
// 	return
// }
