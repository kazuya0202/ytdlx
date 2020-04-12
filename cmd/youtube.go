package cmd

import (
	"os/exec"
	"regexp"

	"github.com/fatih/color"
)

// Youtube is struct
type Youtube struct {
	URL string
	ID  string
	Pat *Patterns
}

// Patterns is struct
type Patterns struct {
	YT  *regexp.Regexp
	ID  *regexp.Regexp
	URL *regexp.Regexp
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
			YT:  yt,
			ID:  id,
			URL: url,
		},
	}
	youtube.extractID()
	return &youtube
}

func (y *Youtube) isAvailable() bool {
	return y.Pat.ID.MatchString(y.ID) && func() bool {
		print(color.YellowString("Validating..."))
		err := exec.Command(cu.CmdName, "-s", y.ID).Run()
		print("\r")

		return err == nil
	}()
}

func (y *Youtube) extractID() {
	if y.isYoutubeURL() {
		y.ID = y.Pat.YT.FindStringSubmatch(y.URL)[1]
	} else if y.Pat.ID.MatchString(y.URL) {
		y.ID = y.URL
	}
}

func (y *Youtube) isURL() bool {
	return y.Pat.URL.MatchString(y.URL)
}

func (y *Youtube) isYoutubeURL() bool {
	return y.isURL() && y.Pat.YT.MatchString(y.URL)
}
