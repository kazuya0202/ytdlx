package cmd

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"os"
	"strings"

	"github.com/kazuya0202/kz"
	"github.com/spf13/cobra"
)

// execute base command.
const baseCommand string = "youtube-dl"

// ArgDefaults is struct.
type ArgDefaults struct {
	IsM4A               bool
	IsMP4               bool
	IsAvailable         bool
	IsSelect            bool
	IsSelectEachFormat  bool
	IsFindFromAvailable bool
	OutputTitle         string
	IsPlaylist          bool
	Extension           string
}

var (
	yts   []*Youtube
	types *SelectTypes

	defs   ArgDefaults
	cu     *CommandUtility
	status *kz.StatusString

	debug string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ytdlx [URL | ID | FILE_PATH]",
	Short: "An application to make youtube-dl easy to use.",
	Long:  `An application to make youtube-dl easy to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := executeMain(cmd, args)
		if err != nil {
			os.Exit(1)
		}
	},
}

func executeMain(cmd *cobra.Command, args []string) error {
	status = kz.NewStatusString()
	types = newSelectTypes() // initialize select download type
	cu = newCommandUtility(baseCommand)

	if len(args) < 1 {
		args = append(args, GetInput("enter"))
	}

	// flatten urls.
	var urls []string
	for _, arg := range args {
		if Exists(arg) {
			urls = append(urls, readFileContent(arg)...)
		} else {
			urls = append(urls, arg)
		}
	}
	// construct
	for _, url := range urls {
		url = strings.TrimSpace(url)
		// ignore commented line or empty line.
		if len(url) > 1 && url[:1] != "#" {
			yts = append(yts, newYoutube(url))
		}
	}

	// if select only once.
	if defs.IsSelect && !defs.IsSelectEachFormat {
		types.selectType()
	}
	cu.determineOption(types)

	// execute youtube-dl
	// batch processing.
	if !(defs.IsSelectEachFormat || defs.IsFindFromAvailable) {
		var array []string
		for _, y := range yts {
			array = append(array, y.URL)
		}
		URLs := strings.Join(array, " ")
		cu.beforeExecute(URLs)
		return cu.execute()
	}

	// each processing.
	for _, y := range yts {
		// if defs.IsSelectEachFormat {
		// 	// select every download.
		// 	types.selectType()
		// 	cu.determineOption(types)
		// } else if defs.IsFindFromAvailable {
		// 	// select every download by using fzf.
		// 	cu.selectAvailableTypes(y.URL)
		// }

		cu.beforeExecute(y.URL)
		if err := cu.execute(); err != nil {
			os.Exit(1)
		}
	}
	return nil
}

func getExtension() string {
	var format string
	if defs.Extension == "" {
		if defs.IsM4A {
			format = "m4a"
		} else {
			format = "mp4"
		}
	} else {
		format = defs.Extension
	}
	return format
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&defs.IsM4A, "audio", "a", false, "Download audio format only")
	rootCmd.Flags().BoolVarP(&defs.IsMP4, "video", "v", false, "Download video format only")
	rootCmd.Flags().BoolVarP(&defs.IsAvailable, "format-list", "F", false, "Show available format list")
	rootCmd.Flags().BoolVarP(&defs.IsSelect, "select", "s", false, "Download selected format")
	// rootCmd.Flags().BoolVarP(&defs.IsSelectEachFormat, "select-each", "S", false, "Download each selected format")
	// rootCmd.Flags().BoolVarP(&defs.IsFindFromAvailable, "find", "f", false, "Download selected from available format list")
	rootCmd.Flags().StringVarP(&defs.OutputTitle, "output", "o", "", "Output filename")
	rootCmd.Flags().BoolVarP(&defs.IsPlaylist, "playlist", "p", false, "Download the playlist (option: --yes-playlist)")
	rootCmd.Flags().StringVarP(&defs.Extension, "ext", "e", "", "Specify download extension (e.g. m4a, mp3, mp4, ogg, webm)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
