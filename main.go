package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

type FileStats struct {
	Path  string
	Lines int
}

func main() {
	fileFlag := flag.String("file", "", "Path to a file")
	dirFlag := flag.String("dir", "", "Path to a directory")
	flag.Parse()

	if (*fileFlag == "" && *dirFlag == "") || (*fileFlag != "" && *dirFlag != "") {
		fmt.Println("Usage: go run linecounter.go [-file <file>] [-dir <directory>]")
		return
	}

	var fileStats []FileStats
	var totalLines int

	if *fileFlag != "" {
		lines, err := countLines(*fileFlag)
		if err != nil {
			log.Fatal(err)
		}
		fileStats = append(fileStats, FileStats{Path: *fileFlag, Lines: lines})
		totalLines += lines
	} else if *dirFlag != "" {
		err := filepath.Walk(*dirFlag, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				return nil
			}

			if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
				lines, err := countLines(path)
				if err != nil {
					log.Println(err)
					return nil
				}
				fileStats = append(fileStats, FileStats{Path: path, Lines: lines})
				totalLines += lines
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Usage: go run linecounter.go [-file <file>] [-dir <directory>]")
		return
	}

	if len(fileStats) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)
		fmt.Fprintf(w, "File\tLines\n")
		fmt.Fprintf(w, "----\t-----\n")
		for _, stat := range fileStats {
			fmt.Fprintf(w, "%s\t%d\n", stat.Path, stat.Lines)
		}
		w.Flush()

		fmt.Printf("\nTotal Lines: %d\n", totalLines)
	} else {
		fmt.Println("No files found.")
	}
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lines := 0
	inMultilineComment := false

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)

		// Ignore empty lines
		if line == "" {
			continue
		}

		// Ignore single-line comments
		if strings.HasPrefix(line, "//") {
			continue
		}

		// Ignore multi-line comments
		if strings.HasPrefix(line, "/*") {
			inMultilineComment = true
		}

		if !inMultilineComment {
			lines++
		}

		// Check if multi-line comment ends
		if strings.HasSuffix(line, "*/") {
			inMultilineComment = false
		}
	}

	return lines, nil
}
