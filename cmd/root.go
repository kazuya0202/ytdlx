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
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// execute base command.
const baseCommand string = "youtube-dl"

// ArgDefaults is struct.
type ArgDefaults struct {
	IsM4A               bool
	IsMP4               bool
	IsBest              bool
	IsDefault           bool
	IsAvailable         bool
	IsSelect            bool
	IsSelectEachFormat  bool
	IsFindFromAvailable bool
	OutputTitle         string

	// Format      string  // TODO
}

var (
	defs ArgDefaults
	cu   CommandUtility
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ytdlx [YouTube-URL | YouTuve-ID]",
	Short: "The command to make youtube-dl easy to use.",
	Long:  `The command to make youtube-dl easy to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		var yts []*Youtube

		cu.setCommandName(baseCommand)
		cu.determineEnvCommand()

		if len(args) < 1 {
			print(color.GreenString("enter") + "> ")
			args = append(args, GetInput())
		}

		// append all target.
		for _, arg := range args {
			if Exists(arg) {
				// if arg is file
				for _, x := range readFileContent(arg) {
					yts = append(yts, newYoutube(x))
				}
			} else {
				// URL or ID.
				yts = append(yts, newYoutube(arg))
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

		for i, y := range yts {
			println("\n>", y.URL)

			if !y.isAvailable() {
				log.Printf("[%s]: '%s' is not valid URL.\n", color.BlueString("ERROR"), y.URL)
				continue
			}

			if defs.IsSelectEachFormat {
				// select every download.
				st.selectType()
				cu.determineOption(st)
			} else if defs.IsFindFromAvailable {
				// select every download by using fzf.
				cu.selectAvailableTypes(y.ID)
			}

			// command
			cu.Arg = fmt.Sprintf("%s %s %s", cu.CmdName, cu.Option, y.ID)

			if defs.OutputTitle != "" {
				if isMulti {
					cu.Arg += fmt.Sprintf(" -o %s_%03d", defs.OutputTitle, i+1)
				} else {
					cu.Arg += fmt.Sprintf(" -o %s", defs.OutputTitle)
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
	rootCmd.Flags().BoolVarP(&defs.IsFindFromAvailable, "find", "f", false, "Downloads selected from available format list")
	rootCmd.Flags().BoolVarP(&defs.IsDefault, "default", "d", false, "Downloads default format")
	rootCmd.Flags().BoolVarP(&defs.IsM4A, "audio", "a", false, "Downloads audio format only")
	rootCmd.Flags().BoolVarP(&defs.IsMP4, "video", "v", false, "Downloads video format only")
	rootCmd.Flags().BoolVarP(&defs.IsBest, "best", "b", false, "Downloads best format")
	rootCmd.Flags().BoolVarP(&defs.IsSelect, "select", "s", false, "Downloads selected format")
	rootCmd.Flags().BoolVarP(&defs.IsSelectEachFormat, "select-each", "S", false, "Downloads each selected format")
	rootCmd.Flags().StringVarP(&defs.OutputTitle, "output", "o", "", "Output filename")
	// rootCmd.Flags().StringVarP(&defs.format, "format", "f", "", "specify format")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
