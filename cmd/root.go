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
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	kz "github.com/kazuya0202/kazuya0202"
)

// ArgDefaults ...
type ArgDefaults struct {
	IsM4A       bool
	IsMP4       bool
	IsFullHD    bool
	IsBest      bool
	IsDefault   bool
	IsAvailable bool
	IsSelect    bool
	// Format      string
	OutputTitle string
}

// CommandConfing ...
type CommandConfing struct {
	cmdName string
	URL     string
	Option  string

	ID string

	IsURL    bool
	IsID     bool
	IsExists bool
}

type selectType struct {
	Default   string
	AudioOnly string
	VideoOnly string
	FullHD    string
	Best      string
	Available string
	Select    string
	// Format    string
	// OutputTitle string

	arrayS   []string
	selected string
	idx      int
}

var (
	defs ArgDefaults
	cu   CommandUtility
)

const ytdlCommand string = "youtube-dl"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ytdlx [YouTube-URL | YouTuve-ID]",
	Short: "The command to make youtube-dl easy to use.",
	Long:  `The command to make youtube-dl easy to use.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var yts []*Youtube

		cu.setCommandName(ytdlCommand)
		cu.determineEnvCommand()

		if len(args) < 1 {
			print(color.GreenString("enter") + "> ")
			args = append(args, kz.GetInput())
		}

		// append all target.
		for _, arg := range args {
			if kz.Exists(arg) {
				// if arg is file
				for _, x := range readFileContent(arg) {
					yts = append(yts, newYoutube(x))
				}
			} else {
				// URL or ID.
				yts = append(yts, newYoutube(arg))
			}
		}

		for _, y := range yts {
			// y.showMessage()
			println("\n>", y.URL)

			if y.isAvailable() {
				cu.execute(y.ID)
			} else {
				log.Printf("[%s]: '%s' is not valid URL.\n", color.BlueString("Warning"), y.URL)
			}
		}

		os.Exit(0)

		var cconf CommandConfing
		cconf.cmdName = ytdlCommand

		// URL
		if len(args) > 0 {
			cconf.URL = args[0]
		}

		cconf.allCheck()

		if !cconf.any() {
			println("Enter [URL | ID] or .txt file path.")

			for !cconf.any() {
				print(color.GreenString("enter") + "> ")
				cconf.URL = kz.GetInput()
				cconf.allCheck()
			}
		}

		log.Printf("[%s]: %s\n", color.BlueString("Processing"), cconf.URL)

		var st selectType

		st.Default = "Default"
		st.AudioOnly = "Audio only"
		st.VideoOnly = "Video only"
		st.FullHD = "Full HD"
		st.Best = "Best"
		st.Select = "Select"
		st.Available = "#Available list"
		// st.Format = "#Format"
		// st.OutputTitle = "Title"
		st.setStringArray()

		st._select()
		cconf.determineOption(&st)
		if st.selected == st.Select || defs.IsSelect {
			cconf.selectOptions()
		}

		// execute URL
		if cconf.IsURL || cconf.IsID {
			cconf.execYtdl()
		}

		// execute URL in text file
		if cconf.IsExists {
			fp, err := os.Open(cconf.URL)
			kz.CheckErr(err)
			defer fp.Close()

			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				cconf.URL = scanner.Text()

				if cconf.checkURL() || cconf.checkID() {
					cconf.execYtdl()
				} else {
					println(cconf.URL, "is not URL.")
				}
				println()
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ytdlx.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&defs.IsAvailable, "format-list", "F", false, "Show available format list")
	rootCmd.PersistentFlags().BoolVarP(&defs.IsDefault, "default", "d", false, "Download default format")
	rootCmd.PersistentFlags().BoolVarP(&defs.IsM4A, "audio", "a", false, "Download audio format only")
	rootCmd.PersistentFlags().BoolVarP(&defs.IsMP4, "video", "v", false, "Download video format only")
	rootCmd.PersistentFlags().BoolVarP(&defs.IsFullHD, "full-hd", "f", false, "Download full HD format")
	rootCmd.PersistentFlags().BoolVarP(&defs.IsBest, "best", "b", false, "Download best format")
	rootCmd.PersistentFlags().BoolVarP(&defs.IsSelect, "select", "s", false, "Download select format")
	rootCmd.PersistentFlags().StringVarP(&defs.OutputTitle, "output", "o", "", "Output filename")
	// rootCmd.PersistentFlags().StringVarP(&defs.format, "format", "f", "", "specify format")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
