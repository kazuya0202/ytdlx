package main

import (
	"regexp"
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

// root
func main() {
	// yts := make(newYoutube("a"), 1)
	xx := []string{"/\\", "b", "c"}
	var yts []*Youtube
	for _, x := range xx {
		yts = append(yts, newYoutube(x))
	}

	for _, y := range yts {
		y.showMessage()
	}
}

func newYoutube(str string) *Youtube {
	// pattern for extracting.
	yt, _ := regexp.Compile(`^.*\/(?:watch\?)?(?:v=)?(?:feature=[a-z_]+&)?(?:v=)?([a-zA-Z0-9-=_]+)(?:&feature=[a-z_]*)?(?:\?t=[0-9]+)?$`)
	// pattern for checking id.
	id, _ := regexp.Compile(`[a-zA-Z0-9-=_]+`) // TODO: スラッシュとかと一緒にアルファベットがあったらそれだけで判定される
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

func (y *Youtube) isOK() bool {
	return y.ID != "" && y.Pat.id.MatchString(y.ID)
}

func (y *Youtube) showMessage() {
	// TODO: else{} <- remove
	if !y.isOK() {
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
