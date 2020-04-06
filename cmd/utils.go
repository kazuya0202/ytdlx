package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strings"

	kz "github.com/kazuya0202/kazuya0202"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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

// ExecCmdInRealTime executes command in real time.
func ExecCmdInRealTime(cmd *exec.Cmd) error {
	stdout, err := cmd.StdoutPipe()
	kz.CheckErr(err)
	stderr, err := cmd.StderrPipe()
	kz.CheckErr(err)

	err = cmd.Start()

	streamReader := func(sc *bufio.Scanner, outChan chan string, doneChan chan bool) {
		defer close(outChan)
		defer close(doneChan)
		for sc.Scan() {
			outChan <- sc.Text()
		}
		doneChan <- true
	}

	stdoutScanner := bufio.NewScanner(stdout)
	stdoutOutputChan := make(chan string)
	stdoutDoneChan := make(chan bool)
	stderrScanner := bufio.NewScanner(stderr)
	stderrOutputChan := make(chan string)
	stderrDoneChan := make(chan bool)
	go streamReader(stdoutScanner, stdoutOutputChan, stdoutDoneChan)
	go streamReader(stderrScanner, stderrOutputChan, stderrDoneChan)

	stillGoing := true
	for stillGoing {
		select {
		case <-stdoutDoneChan:
			stillGoing = false
		case line := <-stdoutOutputChan:
			fmt.Println(SjisToUtf8(line))
		case line := <-stderrOutputChan:
			fmt.Println(SjisToUtf8(line))
		}
	}
	ret := cmd.Wait()
	if ret != nil {
		log.Fatal(err)
	}

	return nil
}

// SjisToUtf8 decodes Shift-JIS to UTF8.
func SjisToUtf8(str string) string {
	iostr := strings.NewReader(str)
	rio := transform.NewReader(iostr, japanese.ShiftJIS.NewDecoder())
	ret, err := ioutil.ReadAll(rio)
	if err != nil {
		return ""
	}
	return string(ret)
}
