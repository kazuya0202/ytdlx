package cmd

import (
	"regexp"
)

// https://www.youtube.com/watch?v=aT_cmdRwxoo
// https://youtu.be/9E8eakpIsNE

// Youtube is struct
type Youtube struct {
	URL     string
	ID      string
	Pattern *regexp.Regexp
}

const expr string = `^.*\/(?:watch\?)?(?:v=)?(?:feature=[a-z_]+&)?(?:v=)?([a-zA-Z0-9-=_]+)(?:&feature=[a-z_]*)?(?:\?t=[0-9]+)?$`

func newYoutube(str string) *Youtube {
	r, _ := regexp.Compile(expr)
	return &Youtube{
		URL:     str,
		ID:      "",
		Pattern: r,
	}
}

func (yt *Youtube) isOK() bool {
	return yt.ID != ""
}

func (yt *Youtube) extractID() {
	if yt.isURL() && yt.isYoutubeURL() {
		yt.ID = yt.Pattern.FindStringSubmatch(yt.URL)[1]
	} else {
		if matched, _ := regexp.MatchString(`[a-zA-Z0-9-=_]+`, yt.URL); matched {
			yt.ID = yt.URL
		}
	}
}

func (yt *Youtube) isURL() bool {
	m, _ := regexp.MatchString(`http(s)?://([\w-]+\.)+[\w-]+(/[\w- ./?%&=]*)?`, yt.URL)
	return m
}

func (yt *Youtube) isYoutubeURL() bool {
	return yt.Pattern.MatchString(yt.URL)
}
