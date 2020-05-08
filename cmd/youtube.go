package cmd

import (
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

	// check executable with sumilating youtube
	err := exec.Command(cu.CmdName, "-s", y.URL, cu.Option).Run()
	return err == nil
}
