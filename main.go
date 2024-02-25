package main

import (
	"flag"
	"fmt"
	cp "github.com/otiai10/copy"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func templates(folder string) ([]string, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	out := make([]string, len(files))

	for i, file := range files {
		out[i] = file.Name()
	}

	return out, nil
}

func createTemplate(name, templateDir, template string) error {
	srcPath := filepath.Join(templateDir, template)
	dstPath := filepath.Join("./", name)

	if _, err := os.Stat(dstPath); err == nil {
		return fmt.Errorf("File '%s' already exists\n", dstPath)
	}

	err := cp.Copy(srcPath, dstPath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var listTemplates bool
	flag.BoolVar(&listTemplates, "l", false, "List all templates")
	flag.Parse()

	templateDir := "/home/rhys/Templates"

	if listTemplates || flag.NArg() < 1 {
		templateArry, err := templates(templateDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, template := range templateArry {
			fmt.Printf("%s\n", template)
		}
	} else {
		requested := flag.Arg(0)
		templateArry, err := templates(templateDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, template := range templateArry {
			match, err := regexp.MatchString(requested, template)
			if err != nil {
				log.Fatal(err)
			}

			if match {
				var outName string
				if flag.NArg() >= 2 {
					outName = flag.Arg(1)
				} else {
					outName = template
				}

				err := createTemplate(outName, templateDir, template)
				if err != nil {
					log.Fatal(err)
				}

				break
			}
		}
	}
}
