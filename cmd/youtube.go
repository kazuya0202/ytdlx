package cmd

import (
	"os/exec"
	"strings"
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

	arg := strings.Join([]string{y.URL, cu.Option}, " ")
	cmd := exec.Command(cu.CmdName, "-s", arg)

	// check executable with sumilating youtube
	return cmd.Run() == nil
}
