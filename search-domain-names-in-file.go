package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	var fp string
	flag.StringVar(&fp, "f", "", "path to file with contents")
	flag.Parse()

	if fp == "" {
		printHelp()
		os.Exit(0)
	}

	data, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalln("Error reading file", fp)
	}

	// https://regex101.com/r/njuQjJ/2
	dRegex := `([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}`

	re := regexp.MustCompile(dRegex)
	for _, domain := range re.FindAll(data, -1) {
		fmt.Println(string(domain))
	}
}

func printHelp() {
	fmt.Println("Usage: -f file path")
	os.Exit(0)
}
