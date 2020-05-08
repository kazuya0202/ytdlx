package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

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

func (c *CommandUtility) setCommand() {
	c.Command = exec.Command(c.EnvCmd.Cmd, c.EnvCmd.Option, c.Arg)
}

func (c *CommandUtility) setCommandName(cmdName string) {
	c.CmdName = cmdName
}

func (c *CommandUtility) shapeCommandString() string {
	// <appName> <opt> <cmdName> ... -> <cmdName> ...
	str := c.Command.String()
	str = str[strings.Index(str, c.CmdName):]
	str = strings.ReplaceAll(str, "  ", " ") // "  " -> " "
	return str
}

func (c *CommandUtility) execute() {
	c.setCommand()
	// println(c.shapeCommandString())
	ExecCmdInRealTime(c.Command)
}

func (c *CommandUtility) determineOption(st selectType) {
	// show available list.
	if defs.IsAvailable || st.isMatched(st.Available) {
		c.Option = "-F"
		return
	}

	// audio download.
	if defs.IsM4A || st.isMatched(st.AudioOnly) {
		// c.Option = "-f 'bestaudio[ext=m4a]/bestaudio'"
		c.Option = "-f bestaudio"
		return
	}

	// video download.
	if defs.IsMP4 || st.isMatched(st.VideoOnly) {
		c.Option = "-f bestvideo[ext=mp4]/bestvideo"
		return
	}

	// select format each URL.
	if st.isMatched(st.SelectEachFormat) {
		defs.IsSelectEachFormat = true
		return
	}

	// find from available list.
	if st.isMatched(st.FindFromAvailable) {
		defs.IsFindFromAvailable = true
		return
	}

	// default download.
	c.Option = "bestvideo+bestaudio/best"
	// if st.isMatched(st.Default) {
	// 	c.Option = "bestvideo+bestaudio/best"
	// 	return
	// }
}

func (c *CommandUtility) selectAvailableTypes(URL string) {
	arg := fmt.Sprintf("%s %s %s", c.CmdName, "-F", URL)
	stdout, _ := exec.Command(c.EnvCmd.Cmd, c.EnvCmd.Option, arg).Output()
	array := strings.Split(string(stdout), "\n")

	var selectable []string
	for i := len(array) - 1; i >= 0; i-- {
		x := array[i]
		if len(x) > 0 {
			if _, err := strconv.Atoi(x[:1]); err == nil {
				selectable = append(selectable, x)
			}
		}
	}

	idxs, _ := fuzzyfinder.FindMulti(
		selectable,
		func(i int) string { return selectable[i] },
	)

	var selected []int
	for _, idx := range idxs {
		t := selectable[idx]
		x := t[:strings.Index(t, " ")]
		if v, err := strconv.Atoi(x); err == nil {
			selected = append(selected, v)
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
}
