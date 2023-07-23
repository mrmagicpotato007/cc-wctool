package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	count      bool
	lines      bool
	wordcount  bool
	charcount bool
	fileroute  bool
	stdinroute bool
)

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("Didnt recive input filename considering stdin")
	}
	fileName := ""
	pwd, _ := os.Getwd()
	for _, arg := range args {
		if arg == "-c" {
			count = true
		}
		if arg == "-l" {
			lines = true
		}
		if arg == "-w" {
			wordcount = true
		}
		if arg == "-m" {
			charcount = true
		}
	}
	if !count && !lines && !wordcount {
		count = true
		lines = true
		wordcount = true
	}
	if len(args) > 1 {
		fileName = args[len(args)-1]
		if fileName != "" && strings.Contains(fileName, ".txt") {
			filePath := filepath.Join(pwd, fileName)
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			stats, err := os.Stat(fileName)
			if err != nil {
				log.Fatal(err)
			}

			if count {
				log.Printf("size of file,%d", stats.Size())
			}
			scanner := bufio.NewScanner(file)
			GenerateWcAndLines(scanner)
			fileroute = true
		} else {
			log.Println("Please input the file name")
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		stdinroute = true
		GenerateWcAndLines(scanner)
	}
	if fileroute && stdinroute {
		log.Println("Input is missing")
	}
}

func GenerateWcAndLines(scanner *bufio.Scanner) {
	wc := 0
	size := 0
	cc := 0
	if lines {
		lines := 0
		for scanner.Scan() {
			if stdinroute {
				size += len(scanner.Bytes())
			}
			line := scanner.Text()
			if line  == "stop reading"{
				break
			}
			cc+=strings.Count(line,"")
	
			x := strings.Split(scanner.Text(), " ")
			wc += len(x)
			lines++
		}
		if count && stdinroute{
			log.Printf("Size of file,%d", size)
		}
		log.Printf("No of lines,%d", lines)
	}
	if charcount{
		log.Printf("No of chars,%d", cc)
	}
	if wordcount {
		log.Printf("No of words,%d", wc)
	}
}
