package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	kz "github.com/kazuya0202/kazuya0202"
	"github.com/ktr0731/go-fuzzyfinder"
)

func (cc *CommandConfing) determineOption(st *selectType) {
	selected := st.selected

	// determine option
	if defs.IsDefault || selected == st.Default {
		// default download
		cc.Option = ""
	} else if defs.IsAvailable || selected == st.Available {
		// format list
		cc.Option = "-F"
	} else if defs.IsM4A || selected == st.AudioOnly {
		// audio download
		cc.Option = "-f bestaudio[ext=m4a]"
	} else if defs.IsMP4 || selected == st.VideoOnly {
		// video download
		cc.Option = "-f bestvideo[ext=mp4]"
	} else if defs.IsBest || selected == st.Best {
		// best format download
		cc.Option = "-f bestvideo[ext=mp4]+bestaudio[ext=m4a] --merge-output-format mp4"
	} else if defs.IsFullHD || selected == st.FullHD {
		// full hd download
		cc.Option = "-f 137+140 --merge-output-format mp4"
	}
}

func (cc *CommandConfing) any() bool {
	return cc.IsURL || cc.IsID || cc.IsExists
}

func (cc *CommandConfing) all() bool {
	return cc.IsURL && cc.IsID && cc.IsExists
}

func (cc *CommandConfing) checkID() bool {
	cmd := exec.Command(cc.cmdName, "-e", cc.URL)
	err := cmd.Run()

	if err == nil {
		cc.IsID = true
	} else {
		cc.IsID = false
	}
	return cc.IsID
}

func (cc *CommandConfing) checkURL() bool {
	matched, err := regexp.MatchString(`http(s)?://([\w-]+\.)+[\w-]+(/[\w- ./?%&=]*)?`, cc.URL)
	kz.CheckErr(err)

	cc.IsURL = matched
	return cc.IsURL
}

func (cc *CommandConfing) checkExists() bool {
	_, err := os.Stat(cc.URL)

	cc.IsExists = (err == nil)
	return cc.IsExists
}

func (cc *CommandConfing) allCheck() (isURL bool, isID bool, isExists bool) {
	x := cc.checkURL()
	y := cc.checkID()
	z := cc.checkExists()
	return x, y, z
}

func (cc *CommandConfing) selectOptions() {
	cc.Option = "-F"
	envCmd := getEnvCommand()
	arg := fmt.Sprintf("%s %s %s", cc.cmdName, cc.Option, cc.URL)

	command := exec.Command(envCmd.Cmd, envCmd.Option, arg)
	stdout, _ := command.Output()
	array := strings.Split(string(stdout), "\n")

	var selectStrings []string
	for i := len(array) - 1; i >= 0; i-- {
		x := array[i]
		if len(x) > 0 {
			if _, err := strconv.Atoi(x[:1]); err == nil {
				selectStrings = append(selectStrings, x)
			}
		}
	}

	idxs, _ := fuzzyfinder.FindMulti(
		selectStrings,
		func(i int) string { return selectStrings[i] },
	)

	var selected []int
	for _, i := range idxs {
		t := selectStrings[i]
		x := t[:strings.Index(t, " ")]
		if i, err := strconv.Atoi(x); err == nil {
			selected = append(selected, i)
		}
	}

	if len(selected) == 0 {
		cc.Option = ""
		println("Download with default format.")
	} else if len(selected) == 1 {
		cc.Option = fmt.Sprintf("-f %d", selected[0])
	} else if len(selected) == 2 {
		// swap
		if selected[0] > selected[1] {
			selected[1], selected[0] = selected[0], selected[1]
		}
		cc.Option = fmt.Sprintf("-f %d+%d", selected[0], selected[1])
		cc.Option += " --merge-output-format mp4"
	} else {
		println("Cannot select more than 3.")
		panic(-1)
	}

	// execute
	cc.execYtdl()
}

func (cc *CommandConfing) execYtdl() {
	envCmd := getEnvCommand()
	cc.appendOutputTitle()
	arg := fmt.Sprintf("%s %s %s", cc.cmdName, cc.Option, cc.URL)

	command := exec.Command(envCmd.Cmd, envCmd.Option, arg)
	log.Printf("[%s]: %s\n", color.BlueString("Execute"), command.String())
	ExecCmdInRealTime(command)
	// os.Exit(1)
}

func (cc *CommandConfing) appendOutputTitle() {
	if defs.OutputTitle != "" {
		cc.Option += fmt.Sprintf(" -o %s", defs.OutputTitle)
		cc.Option = strings.TrimSpace(cc.Option)
	}
}
