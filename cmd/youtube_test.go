package cmd

import (
	"fmt"
	"os"
	"testing"

	kz "github.com/kazuya0202/kazuya0202"
)

func TestYoutube(t *testing.T) {
	// xx := []string{"com/\\", "b", "c"}
	var yts []*Youtube
	// // fmt.Println(xx)
	// // fmt.Println(yts)
	// for _, x := range xx {
	// 	yts = append(yts, newYoutube(x))
	// }

	// for _, y := range yts {
	// 	y.showMessage()
	// }

	var cc CommandConfing
	cc.cmdName = ytdlCommand

	args := os.Args[3:]
	fmt.Println(args)
	fmt.Println(len(args))

	if len(args) < 3 {
		input := kz.GetInput()
		args = append(args, input)
	}

	for _, arg := range args {
		yts = append(yts, newYoutube(arg))
	}
	yts = yts[:len(yts)-1] // adjust

	for _, y := range yts {
		// y.showMessage()
		if y.isAvailable() {
			println("kore ID.")
			y.execYtdl(&cc)
		} else {
			println("kore not ID.")
		}
	}
}
