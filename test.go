package main

import (
	"regexp"
)

func main() {
	r, _ := regexp.Compile(`^.*\/(?:watch\?)?(?:v=)?(?:feature=[a-z_]+&)?(?:v=)?([a-zA-Z0-9-=_]+)(?:&feature=[a-z_]*)?(?:\?t=[0-9]+)?$`)
	url := "https://www.youtube.com/watch?v=aT_cmdRwxoo"
	url2 := "https://youtu.be/9E8eakpIsNE"

	// if isurl(url) {
	// fmt.Println(r.FindStringSubmatch(url)[1])
	// }
	// fmt.Println(r.FindStringSubmatch(url2)[1])
	// fmt.Println(r.FindStringSubmatch("aT_cmdRwxoo"))

	x := r.MatchString(url)
	y := r.MatchString(url2)
	z := r.MatchString("aaa")
	println(x, y, z)
}
