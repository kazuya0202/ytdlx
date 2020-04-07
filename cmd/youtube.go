package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// https://www.youtube.com/watch?v=aT_cmdRwxoo
// https://youtu.be/9E8eakpIsNE

// Youtube is struct
type Youtube struct {
	URL string
	ID  string
	Pat *Patterns
}

// Patterns is struct
type Patterns struct {
	yt  *regexp.Regexp
	id  *regexp.Regexp
	url *regexp.Regexp
}

// newYoutube creates initialized Youtube instance.
func newYoutube(str string) *Youtube {
	// pattern for extracting.
	yt, _ := regexp.Compile(`^.*\/(?:watch\?)?(?:v=)?(?:feature=[a-z_]+&)?(?:v=)?([a-zA-Z0-9-=_]+)(?:&feature=[a-z_]*)?(?:\?t=[0-9]+)?$`)
	// pattern for checking id.
	id, _ := regexp.Compile(`^[a-zA-Z0-9-=_]+$`)
	// pattern for checking url.
	url, _ := regexp.Compile(`http(s)?://([\w-]+\.)+[\w-]+(/[\w- ./?%&=]*)?`)

	youtube := Youtube{
		URL: str,
		Pat: &Patterns{
			yt:  yt,
			id:  id,
			url: url,
		},
	}
	youtube.extractID()
	return &youtube
}

func (y *Youtube) isAvailable() bool {
	return y.ID != "" && y.Pat.id.MatchString(y.ID)
}

func (y *Youtube) showMessage() {
	// TODO: else{} <- remove
	if !y.isAvailable() {
		println(y.URL, "is not correct.")
	} else {
		println("[DEBUG]:", y.URL, "is correct.")
		println("id:", y.ID)
	}
}

func (y *Youtube) extractID() {
	if y.isYoutubeURL() {
		y.ID = y.Pat.yt.FindStringSubmatch(y.URL)[1]
	} else if y.Pat.id.MatchString(y.URL) {
		y.ID = y.URL
	}
}

func (y *Youtube) isURL() bool {
	return y.Pat.url.MatchString(y.URL)
}

func (y *Youtube) isYoutubeURL() bool {
	return y.isURL() && y.Pat.yt.MatchString(y.URL)
}

func (y *Youtube) execYtdl(cu *CommandUtility) {
	// get environment command
	envCmd := getEnvCommand()

	// otuput option
	if defs.OutputTitle != "" {
		cu.Option += fmt.Sprintf(" -o %s", defs.OutputTitle)
		cu.Option = strings.TrimSpace(cu.Option)
	}
	// y.appendOutputTitle()

	arg := fmt.Sprintf("%s %s %s", cu.CmdName, cu.Option, y.URL)

	command := exec.Command(envCmd.Cmd, envCmd.Option, arg)
	log.Printf("[%s]: %s\n", color.BlueString("Execute"), command.String())
	// ExecCmdInRealTime(command)
	// os.Exit(1)
}
