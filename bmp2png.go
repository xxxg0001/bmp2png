package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
)

func Walkfunc(path string, info os.FileInfo, err error) error {
	if filepath.Ext(info.Name()) == ".bmp" {

		name := strings.TrimSuffix(info.Name(), ".bmp")
		dir := filepath.Dir(path)
		fmt.Println(filepath.Join(dir, name+".png"))
		f1, err := os.Open(path)
		img, err := bmp.Decode(f1)
		f2, err := os.Create(filepath.Join(dir, name+".png"))

		png.Encode(f2, img)

		return err
	}
	return nil
}
func main() {
	var path string
	flag.StringVar(&path, "path", "", `C:\bmp`)
	flag.Parse();
	if path == "" {
		fmt.Println("path is null")
		return
	}
	
	filepath.Walk(path, Walkfunc)

}
