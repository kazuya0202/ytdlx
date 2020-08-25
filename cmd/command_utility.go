package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kazuya0202/kz"
)

// CommandUtility is struct.
type CommandUtility struct {
	CmdName string
	Option  string
	Arg     string
	Command *exec.Cmd
	EnvCmd  kz.EnvCommand
}

func newCommandUtility(cmdName string) *CommandUtility {
	var cu CommandUtility
	cu.setCommandName(cmdName)
	cu.determineEnvCommand()

	return &cu
}

func (c *CommandUtility) determineEnvCommand() {
	c.EnvCmd.Determines()
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

func (c *CommandUtility) beforeExecute(URL string) {
	var ytdlOption string
	if strings.Contains(URL, "list=") {
		if defs.IsPlaylist {
			// -ci (--continue --ignore-errors)
			ytdlOption = "--yes-playlist -ci"
		} else {
			ytdlOption = "--no-playlist"
		}
	} else if strings.Contains(URL, " ") {
		ytdlOption += " -ci"
	}

	c.appendOption(ytdlOption)
	c.Arg = fmt.Sprintf("%s %s %s", cu.CmdName, cu.Option, URL)

	// [NAME]-[ID].ext
	// outputTemplate := "%(title)s-%(id)s.%(ext)s"
	if defs.OutputTitle != "" {
		s := fmt.Sprintf("-o %s", defs.OutputTitle+"-%(id)s.%(ext)s")
		c.appendArg(s)
	}

	// for testing.
	// c.Arg = fmt.Sprintf("%s -s %s %s -o '%s'", cu.CmdName, cu.Option, URL, outputTemplate)
	// c.Arg = fmt.Sprintf("%s --get-filename %s %s -o '%s'", cu.CmdName, cu.Option, URL, outputTemplate)
	c.setCommand()
}

func (c *CommandUtility) execute() error {
	// println(c.Command.String())
	println(c.shapeCommandString(), "\n")
	return kz.ExecCmdInRealTime(c.Command)
	// return status.ErrNormal  # for debug.
}

func (c *CommandUtility) clearOption() {
	c.Option = ""
}

func (c *CommandUtility) appendOption(s string) {
	c.Option = strings.Join([]string{c.Option, s}, " ")
}

func (c *CommandUtility) appendArg(s string) {
	c.Arg = strings.Join([]string{c.Arg, s}, " ")
}

func (c *CommandUtility) determineOption(st *SelectTypes) {
	c.clearOption()

	// format option
	c.determineFormatOption(st)
}

func (c *CommandUtility) determineFormatOption(st *SelectTypes) {
	if defs.IsAvailable || st.isMatched(st.Available) {
		// show available list.
		c.appendOption("-F")

	} else if defs.IsM4A || st.isMatched(st.AudioOnly) {
		// audio download.
		s := fmt.Sprintf("-f bestaudio[ext=%s]/bestaudio", getExtension())
		c.appendOption(s)

	} else if defs.IsMP4 || st.isMatched(st.VideoOnly) {
		// video download.
		s := fmt.Sprintf("-f bestvideo[ext=%s]/bestvideo", getExtension())
		c.appendOption(s)

	} else if st.isMatched(st.SelectEachFormat) {
		// select format each URL.
		defs.IsSelectEachFormat = true

	} else if st.isMatched(st.FindFromAvailable) {
		// find from available list.
		defs.IsFindFromAvailable = true

	} else {
		// default download.
		c.appendOption("-f bestvideo+bestaudio/best")

		s := fmt.Sprintf("--merge-output-format %s", getExtension())
		c.appendOption(s)
	}
}

// func (c *CommandUtility) selectAvailableTypes(URL string) {
// 	arg := fmt.Sprintf("%s %s %s", c.CmdName, "-F", URL)
// 	stdout, _ := exec.Command(c.EnvCmd.Cmd, c.EnvCmd.Option, arg).Output()
// 	array := strings.Split(string(stdout), "\n")

// 	var selectable []string
// 	for i := len(array) - 1; i >= 0; i-- {
// 		x := array[i]
// 		if len(x) > 0 {
// 			if _, err := strconv.Atoi(x[:1]); err == nil {
// 				selectable = append(selectable, x)
// 			}
// 		}
// 	}

// 	idxs, _ := fuzzyfinder.FindMulti(
// 		selectable,
// 		func(i int) string { return selectable[i] },
// 	)

// 	var selected []int
// 	for _, idx := range idxs {
// 		t := selectable[idx]
// 		x := t[:strings.Index(t, " ")]
// 		if v, err := strconv.Atoi(x); err == nil {
// 			selected = append(selected, v)
// 		}
// 	}

// 	c.clearOption() // clear option
// 	if len(selected) == 0 {
// 		c.appendOption("-f bestvideo+bestaudio/best")
// 		println("Download with default format.")
// 	} else if len(selected) == 1 {
// 		// c.Option = fmt.Sprintf("-f %d", selected[0])
// 		c.appendOption(fmt.Sprintf("-f %d", selected[0]))
// 	} else if len(selected) == 2 {
// 		// swap
// 		if selected[0] > selected[1] {
// 			selected[1], selected[0] = selected[0], selected[1]
// 		}
// 		c.appendOption(fmt.Sprintf("-f %d+%d", selected[0], selected[1]))

// 		s := fmt.Sprintf("--merge-output-format %s", getExtension())
// 		c.appendOption(s)
// 	} else {
// 		println("Cannot select more than 3.")
// 		panic(-1)
// 	}
// }
