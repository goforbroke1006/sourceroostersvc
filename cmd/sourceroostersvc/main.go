package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/martinlindhe/notify"
)

const ServiceName = "sourceroostersvc"

type Parameters struct {
	Root struct {
		Watch struct {
			Directories []string `yaml:"directories"`
			Files       []string `yaml:"files"`
		} `yaml:"watch"`
	} `yaml:"parameters"`
}

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)

	data, err := ioutil.ReadFile("./parameters.yml")
	checkErr(err)

	p := Parameters{}

	err = yaml.Unmarshal(data, &p)
	checkErr(err)

	dirsText := ""
	for _, dirPath := range p.Root.Watch.Directories {
		fi, err := ioutil.ReadDir(dirPath)
		checkErr(err)
		for _, sub := range fi {
			dirsText = dirsText + sub.Name() + "\n"
		}
	}
	notify.Notify("app name",
		"I watch on your ungly projects, piece of trash!",
		dirsText,
		dir+"/icon.png",
	)

	done := make(chan bool, len(p.Root.Watch.Directories))

	files := make(chan string, 1000)
	go func() {
		for _, dirPath := range p.Root.Watch.Directories {
			findFiles(dirPath, p.Root.Watch.Files, files, done)
		}
	}()

	go func() {
		for f := range files {
			fmt.Println(f)
		}
	}()

	for range p.Root.Watch.Directories {
		<-done
	}

	log.Fatalln("See you!")
}

func findFiles(parentDir string, extList []string, files chan string, done chan bool) {
	go func() {
		filepath.Walk(parentDir, func(path string, f os.FileInfo, _ error) error {
			if !f.IsDir() {
				for _, ext := range extList {
					if ok, _ := regexp.MatchString(ext, f.Name()); ok {
						ap, err := filepath.Abs(path)
						checkErr(err)
						files <- ap
					}
				}
			}
			return nil
		})
		done <- true
	}()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
}
