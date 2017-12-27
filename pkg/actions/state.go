package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

// StateDump returns a function that will log the current app state to the configured logger's output
func StateDump(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		js, err := s.JSONString()
		if err != nil {
			return err
		}
		k.Debugln(js)
		return nil
	}
}
