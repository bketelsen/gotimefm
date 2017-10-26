package grifts

import (
  "github.com/gobuffalo/buffalo"
	"github.com/bketelsen/gotimefm/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
