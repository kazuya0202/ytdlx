package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

// CommandUtility is struct.
type CommandUtility struct {
	CmdName string
	Option  string
	Arg     string
	Command *exec.Cmd
	EnvCmd  envCommand
}

func (c *CommandUtility) setEnvCommand() {
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

	log.Printf("[%s]: %s\n", color.BlueString("Execute"), c.Command.String())
	// ExecCmdInRealTime(c.Command)
}
