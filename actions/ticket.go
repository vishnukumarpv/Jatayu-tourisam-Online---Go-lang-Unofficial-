package actions

import (
	"github.com/gobuffalo/buffalo"
)

func TicketHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("ticket.html"))
}
