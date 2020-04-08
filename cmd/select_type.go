package cmd

import (
	kz "github.com/kazuya0202/kazuya0202"
	"github.com/ktr0731/go-fuzzyfinder"
)

func (st *selectType) setStringArray() {
	st.arrayS = []string{
		st.Default,
		st.AudioOnly,
		st.VideoOnly,
		st.FullHD,
		st.Best,
		st.Available,
		// st.Select,
		// st.Format,
	}
}

func (st *selectType) _select() string {
	// any := false
	// any = any || defs.IsAvailable || defs.IsBest || defs.IsDefault
	// any = any || defs.IsFullHD || defs.IsM4A || defs.IsMP4 || defs.IsSelect

	// if !any {
	if defs.IsSelect {
		// fzf
		idx, err := fuzzyfinder.Find(
			st.arrayS,
			func(i int) string { return st.arrayS[i] },
		)
		kz.CheckErr(err)

		st.selected = st.arrayS[idx]
		println("selected:", st.selected)
	}
	return st.selected
}

func (st *selectType) isMatched(target string) bool {
	return st.selected == target
}
