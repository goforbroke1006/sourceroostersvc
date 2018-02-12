package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	//"regexp"

	"github.com/goforbroke1006/sourceroostersvc"
	"github.com/martinlindhe/notify"
	"time"
	"regexp"
)

const ServiceName = "sourceroostersvc"

type Parameters struct {
	Root struct {
		Watch struct {
			Directories []string `yaml:"directories"`
			Matches     Matches  `yaml:"matches"`
		} `yaml:"watch"`
	} `yaml:"parameters"`
}

type Matches struct {
	Whitelist []string `yaml:"whitelist"`
	Blacklist []string `yaml:"blacklist"`
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
			findFiles(dirPath, p.Root.Watch.Matches, files, done)
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
	time.Sleep(5 * time.Second)

	log.Fatalln("See you!")
}

func findFiles(parentDir string, extList Matches, files chan string, done chan bool) {
	go func() {
		filepath.Walk(parentDir, func(path string, f os.FileInfo, _ error) error {

			for _, black := range extList.Blacklist{
				if ok, _ := regexp.MatchString(black, path); ok {
					return nil
				}
			}

			if !f.IsDir() {
				/*
				for _, ext := range extList.Whitelist {
					if ok, _ := regexp.MatchString(ext, path); ok {
						ap, err := filepath.Abs(path)
						checkErr(err)
						files <- ap
					}
				}*/
			} else {
				if sourceroostersvc.IsProjectDir(path) {
					files <- path
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
