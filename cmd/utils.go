package cmd

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strings"
)

// envCommand is struct.
type envCommand struct {
	Cmd    string // execute command
	Option string // command option
}

// getEnvCommand determines command depend in os environment.
func getEnvCommand() envCommand {
	// windows
	if runtime.GOOS == "windows" {
		return envCommand{"cmd", "/c"}
	}
	// other than windows
	return envCommand{"sh", "-c"}
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
		x := strings.Split(scanner.Text(), " ")[0]
		array = append(array, x)
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
func GetInput() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()
}
