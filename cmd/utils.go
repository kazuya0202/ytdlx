package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/kazuya0202/kz"
)

func escapeAmpersand(URL string) string {
	// windows only | & => ^&
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(URL, "&", "^&")
	}
	return URL
}

// getEnvCommand determines command depend in os environment.
func getEnvCommand() *kz.EnvCommand {
	// windows
	if runtime.GOOS == "windows" {
		return &kz.EnvCommand{
			Cmd:    "cmd",
			Option: "/c",
		}
	}
	// other than windows
	return &kz.EnvCommand{
		Cmd:    "sh",
		Option: "-c",
	}
}

// readFileContent ...
func readFileContent(path string) []string {
	var array []string

	fp, err := os.Open(path)
	CheckErr(err)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		// array = append(array, scanner.Text())
		x := strings.Split(scanner.Text(), "\n")
		for _, line := range x {
			for _, s := range strings.Split(line, " ") {
				if s != "#" {
					array = append(array, s)
				} else {
					break
				}
			}
		}
		// array = append(array, x...)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return array
}

// CheckErr checks nil.
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Exists checks path.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetInput scans input string.
func GetInput(title string) string {
	fmt.Printf("%s> ", color.GreenString(title))

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()
}
