package main

import (
	"flag"
	"fmt"
	"github.com/gographics/imagick/imagick"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// set args
	flag.Parse()
	args := flag.Args()

	// set directories
	readdir, writedir := getRWdir(args)
	pdfFiles := getPdfLists(readdir)
	os.Mkdir(writedir, 0777)

	// convert pdf to img
	for _, filename := range pdfFiles {
		fmt.Println(filename)
		filename := filename + "[0]"
		imgname := writedir + "/" + strings.Replace(filename, "pdf[0]", "png", 1)
		if _, err := os.Stat(imgname); os.IsNotExist(err) {
			fmt.Println("processing...")
			convertPdfToImg(readdir+"/"+filename, imgname)
		}
	}
}

func getRWdir(args []string) (string, string) {
	r := "./"
	w := "./img"
	if len(args) > 0 {
		r = args[0]
	}
	if len(args) > 1 {
		w = args[1]
	}
	return r, w
}

func getPdfLists(readdir string) []string {
	files, _ := ioutil.ReadDir(readdir)
	filenames := []string{}
	for _, f := range files {
		filename := f.Name()
		if strings.HasSuffix(filename, ".pdf") {
			filenames = append(filenames, filename)
		}
	}
	return filenames
}

func convertPdfToImg(filename string, imgname string) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	re := mw.ReadImage(filename)
	if re != nil {
		log.Fatal(re)
	}
	defer mw.Destroy()

	mwc := mw.Clone()

	// mwc.AdaptiveResizeImage(30, 30)
	fmt.Println("png file created: " + imgname)
	w := mwc.WriteImage(imgname)
	if w != nil {
		log.Fatal(w)
	}
}
