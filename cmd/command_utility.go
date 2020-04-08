package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/ktr0731/go-fuzzyfinder"
)

// CommandUtility is struct.
type CommandUtility struct {
	CmdName string
	Option  string
	Arg     string
	Command *exec.Cmd
	EnvCmd  envCommand
}

func (c *CommandUtility) determineEnvCommand() {
	if runtime.GOOS == "windows" {
		// windows
		c.EnvCmd = envCommand{"cmd", "/c"}
	} else {
		// other than windows
		c.EnvCmd = envCommand{"sh", "-c"}
	}
}

func (c *CommandUtility) setCommandName(cmdName string) {
	c.CmdName = cmdName
}

func (c *CommandUtility) appendOutputTitle() {
	if defs.OutputTitle != "" {
		c.Option += fmt.Sprintf(" -o %s", defs.OutputTitle)
		c.Option = strings.TrimSpace(c.Option)
	}
}

func (c *CommandUtility) execute(id string) {
	c.appendOutputTitle()

	c.Arg = fmt.Sprintf("%s %s %s", c.CmdName, c.Option, id)
	c.Command = exec.Command(c.EnvCmd.Cmd, c.EnvCmd.Option, c.Arg)

	log.Printf("[%s]: %s\n", color.BlueString("Running"), c.Command.String())
	// ExecCmdInRealTime(c.Command)
}

func (c *CommandUtility) determineOption(st selectType) {
	// determine option
	if defs.IsDefault || st.isMatched(st.Default) {
		// default download
		c.Option = ""
	} else if defs.IsAvailable || st.isMatched(st.Available) {
		// format list
		c.Option = "-F"
	} else if defs.IsM4A || st.isMatched(st.AudioOnly) {
		// audio download
		c.Option = "-f bestaudio[ext=m4a]"
	} else if defs.IsMP4 || st.isMatched(st.VideoOnly) {
		// video download
		c.Option = "-f bestvideo[ext=mp4]"
	} else if defs.IsBest || st.isMatched(st.Best) {
		// best format download
		c.Option = "-f bestvideo[ext=mp4]+bestaudio[ext=m4a] --merge-output-format mp4"
	} else if defs.IsFullHD || st.isMatched(st.FullHD) {
		// full hd download
		c.Option = "-f 137+140 --merge-output-format mp4"
	}
}

func (c *CommandUtility) selectOptions(id string) {
	c.Option = "-F"

	arg := fmt.Sprintf("%s %s %s", c.CmdName, c.Option, id)
	stdout, _ := exec.Command(c.EnvCmd.Cmd, c.EnvCmd.Option, arg).Output()

	array := strings.Split(string(stdout), "\n")

	var selectStrings []string
	for i := len(array) - 1; i >= 0; i-- {
		x := array[i]
		if len(x) > 0 {
			if _, err := strconv.Atoi(x[:1]); err == nil {
				selectStrings = append(selectStrings, x)
			}
		}
	}

	idxs, _ := fuzzyfinder.FindMulti(
		selectStrings,
		func(i int) string { return selectStrings[i] },
	)

	var selected []int
	for _, i := range idxs {
		t := selectStrings[i]
		x := t[:strings.Index(t, " ")]
		if i, err := strconv.Atoi(x); err == nil {
			selected = append(selected, i)
		}
	}

	if len(selected) == 0 {
		c.Option = ""
		println("Download with default format.")
	} else if len(selected) == 1 {
		c.Option = fmt.Sprintf("-f %d", selected[0])
	} else if len(selected) == 2 {
		// swap
		if selected[0] > selected[1] {
			selected[1], selected[0] = selected[0], selected[1]
		}
		c.Option = fmt.Sprintf("-f %d+%d", selected[0], selected[1])
		c.Option += " --merge-output-format mp4"
	} else {
		println("Cannot select more than 3.")
		panic(-1)
	}

	// execute
	c.execute(id)
}
