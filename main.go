package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("1 argument should be given.")
	}
	path := os.Args[1]
	ignorePath := filepath.Clean(path)
	if filepath.Base(ignorePath) != ".gitignore" {
		log.Fatal("arg should specify .gitignore file.")
	}

	f, err := os.Open(ignorePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	count := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		str := strings.TrimSpace(s.Text())
		if isLineValid(str) {
			err := exec.Command("git", "rm", "--cached", str).Run()
			if err != nil {
				fmt.Println(err)
			} else {
				count++
			}
		}

	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	fmt.Println("REGNORE: " + strconv.Itoa(count) + " files are ignored.")
}

// ignoreのlineが有効かチェック
func isLineValid(line string) bool {
	if line == "" {
		return false
	}
	return !strings.HasPrefix(line, "#")
}
