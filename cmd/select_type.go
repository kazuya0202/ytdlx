package cmd

import (
	"github.com/ktr0731/go-fuzzyfinder"
)

type selectType struct {
	Default           string
	AudioOnly         string
	VideoOnly         string
	Available         string
	SelectEachFormat  string
	FindFromAvailable string

	// Select            string
	// Format    string  // TODO
	// OutputTitle string  // unnecessary

	List     []string
	Selected string
}

func (s *selectType) setStringArray() {
	s.Default = "Default (best)"
	s.AudioOnly = "Audio only"
	s.VideoOnly = "Video only"
	s.SelectEachFormat = "Select each format"
	s.FindFromAvailable = "Find (Select each format from available list)"
	s.Available = "Show available list [* No download]"

	// st.Format = "#Format"  // TODO

	s.List = []string{
		s.Default,
		s.AudioOnly,
		s.VideoOnly,
		s.SelectEachFormat,
		s.FindFromAvailable,
		s.Available,

		// st.Format,
	}
}

func (s *selectType) selectType() {
	if defs.IsSelect || defs.IsSelectEachFormat {
		// fzf
		idx, err := fuzzyfinder.Find(
			s.List,
			func(i int) string { return s.List[i] },
		)
		CheckErr(err)

		s.Selected = s.List[idx]
		println("selected:", s.Selected)
	}
}

func (s *selectType) isMatched(target string) bool {
	return s.Selected == target
}
