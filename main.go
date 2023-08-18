package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	versionFlag bool
	verboseFlag bool
	version     string
	yellow      = color.New(color.FgYellow)
	red         = color.New(color.FgRed)
	blue        = color.New(color.FgBlue)
	green       = color.New(color.FgGreen)
	bold        = color.New(color.FgHiWhite, color.Bold)
)

func init() {
	flag.BoolVar(&versionFlag, "version", false, "Show version")
	flag.BoolVar(&verboseFlag, "V", false, "Verbose output")
	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Printf("\n     📄  nomoreDS %v\n\n", version)
		return
	}
	finder()
}

func pwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path + `/`
}

func finder() {
	var files []string
	err := filepath.Walk(pwd(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".DS_Store" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	remover(files)
}

func remover(files []string) {
	for _, file := range files {
		if verboseFlag {
			fmt.Printf("%s Removing %s...\n", green.Sprintf("[+]"), yellow.Sprintf(file))
		}
		e := os.Remove(file)
		if e != nil {
			if verboseFlag {
				fmt.Printf("%s Error removing %s\n", red.Sprintf("[-]"), yellow.Sprintf(file))
			}
		}
	}
	if verboseFlag {
		fmt.Printf("%s Files deleted: %v\n", blue.Sprintf("[i]"), bold.Sprintf("%d", len(files)))
	}
}
