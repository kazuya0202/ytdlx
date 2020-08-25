package cmd

// Youtube ...
type Youtube struct {
	URL string
}

func newYoutube(str string) *Youtube {
	str = escapeAmpersand(str)
	return &Youtube{URL: str}
}

// func (y *Youtube) isAvailable() bool {
// 	print("Validating...")

// 	if y.URL == "" {
// 		return false
// 	}

// 	// --simulate | youtube-dl -s [OPT] URL
// 	arg := fmt.Sprintf("%s -s %s %s", cu.CmdName, cu.Option, y.URL)
// 	cmd := exec.Command(cu.EnvCmd.Cmd, cu.EnvCmd.Option, arg)
// 	out, err := cmd.Output()

// 	if err != nil {
// 		fmt.Println(string(out))
// 		return false
// 	}

// 	// check executable with sumilating youtube
// 	// return cmd.Run() == nil
// 	return true
// }

// func (y *Youtube) excludePlaylistString() {
// 	idx := strings.Index(y.URL, "list=")
// 	if idx == -1 {
// 		return
// 	}
// 	y.URL = y.URL[:idx-1]
// }
