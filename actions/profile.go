package actions

import (
	"github.com/gobuffalo/buffalo"
)

//ProfileHandler is a
func ProfileHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("profile.html"))
}
