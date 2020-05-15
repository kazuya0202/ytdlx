package cmd

import (
	"fmt"
	"os/exec"
)

// Youtube ...
type Youtube struct {
	URL string
}

func newYoutubeAlpha(str string) *Youtube {
	return &Youtube{URL: str}
}

func newYoutubeAlphaArray(strs []string) []*Youtube {
	var ar []*Youtube
	for _, x := range strs {
		ar = append(ar, newYoutubeAlpha(x))
	}
	return ar
}

func (y *Youtube) isAvailable() bool {
	print("Validating...")

	if y.URL == "" {
		return false
	}

	// --simulate | youtube-dl -s [OPT] URL
	arg := fmt.Sprintf("%s -s %s %s", cu.CmdName, cu.Option, y.URL)
	cmd := exec.Command(cu.EnvCmd.Cmd, cu.EnvCmd.Option, arg)

	// check executable with sumilating youtube
	return cmd.Run() == nil
}
