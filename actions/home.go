package actions

import (
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	// sess := c.Session().Get("current_user_id")
	// c.Flash().Add("success", fmt.Sprintf("logged in %s", sess))
	return c.Render(200, r.HTML("index.html"))
}
