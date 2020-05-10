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

	"github.com/fatih/color"
	kz "github.com/kazuya0202/kazuya0202"
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

	// Format      string  // TODO
}

var (
	defs   ArgDefaults
	cu     CommandUtility
	status *kz.StatusString
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ytdlx [URL | ID]",
	Short: "The command to make youtube-dl easy to use.",
	Long:  `The command to make youtube-dl easy to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		var yts []*Youtube
		status = kz.NewStatusString()

		cu.setCommandName(baseCommand)
		cu.determineEnvCommand()

		if len(args) < 1 {
			args = append(args, kz.GetInput("enter"))
		}

		// append all target.
		for _, arg := range args {
			if Exists(arg) {
				// if arg is file
				ar := readFileContent(arg)
				yts = append(yts, newYoutubeAlphaArray(ar)...)
			} else {
				// URL or ID.
				yts = append(yts, newYoutubeAlpha(arg))
			}
		}

		// select download type
		var st selectType
		st.setStringArray()

		if defs.IsSelect && !defs.IsSelectEachFormat {
			st.selectType()
		}
		cu.determineOption(st)

		// is multi download
		isMulti := len(yts) > 1

		i := 0
		for _, y := range yts {
			i++ // filename index
			println("\n>", y.URL)

			if !y.isAvailable() {
				println(color.RedString("ERROR"))
				s := fmt.Sprintf("'%s' is not valid URL. Skip this URL.", y.URL)
				status.DisplayWarning(s)
				i--
				continue
			}

			println(color.BlueString("SUCCESS"))

			if defs.IsSelectEachFormat {
				// select every download.
				st.selectType()
				cu.determineOption(st)
			} else if defs.IsFindFromAvailable {
				// select every download by using fzf.
				cu.selectAvailableTypes(y.URL)
			}

			// command
			cu.Arg = fmt.Sprintf("%s %s %s", cu.CmdName, cu.Option, y.URL)

			if defs.OutputTitle != "" {
				if isMulti {
					cu.appendArg(fmt.Sprintf("-o %s_%03d", defs.OutputTitle, i))
				} else {
					cu.appendArg(fmt.Sprintf("-o %s", defs.OutputTitle))
				}
			}

			// execute
			cu.execute()
		}
	},
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

	rootCmd.Flags().BoolVarP(&defs.IsAvailable, "format-list", "F", false, "Show available format list")
	rootCmd.Flags().BoolVarP(&defs.IsFindFromAvailable, "find", "f", false, "Download selected from available format list")
	rootCmd.Flags().BoolVarP(&defs.IsM4A, "audio", "a", false, "Download audio format only")
	rootCmd.Flags().BoolVarP(&defs.IsMP4, "video", "v", false, "Download video format only")
	rootCmd.Flags().BoolVarP(&defs.IsSelect, "select", "s", false, "Download selected format")
	rootCmd.Flags().BoolVarP(&defs.IsSelectEachFormat, "select-each", "S", false, "Download each selected format")
	rootCmd.Flags().StringVarP(&defs.OutputTitle, "output", "o", "", "Output filename")
	rootCmd.Flags().BoolVarP(&defs.IsPlaylist, "playlist", "p", false, "Download the playlist (option: --yes-playlist)")

	// rootCmd.Flags().StringVarP(&defs.format, "format", "f", "", "specify format")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
