package actions

import (
	"fmt"
	"projects/gojatayu/models"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// BookingHandler default implementation.
func BookingHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("booking/book.html"))
}

type Total struct {
	adults, children, infants, total float32
}

type Prices struct {
	adults, children, infants float32
}

type Pricing struct {
	jatayu, adventure, heliride Prices
}

// BookHandler default implementation.
func BookHandler(c buffalo.Context) error {
	b := &models.Booking{}
	var pricing = map[string]Prices{
		"jatayu":    {adults: 450.0, children: 400.0, infants: 0.0},
		"adventure": {adults: 1000.0, children: 850.0, infants: -1.0},
		"heliride":  {adults: 2000.0, children: 1800.0, infants: 500.0},
	}
	var t Total
	if err := c.Bind(b); err != nil {
		return errors.WithStack(err)
	}
	b.Adults = 2
	b.Children = 0
	b.Infants = 0
	b.BookedTo = time.Now().Local()
	b.Package = "jatayu"

	packageJatayu := func() Total {
		t.adults = float32(b.Adults) * pricing["jatayu"].adults
		t.children = float32(b.Children) * pricing["jatayu"].children
		t.infants = float32(b.Infants) * pricing["jatayu"].infants
		return t
	}
	packageAdventure := func() Total {
		t.adults = float32(b.Adults) * pricing["adventure"].adults
		t.children = float32(b.Children) * pricing["adventure"].children
		t.infants = float32(b.Infants) * pricing["adventure"].infants
		return t
	}
	packageHeliride := func() Total {
		t.adults = float32(b.Adults) * pricing["heliride"].adults
		t.children = float32(b.Children) * pricing["heliride"].children
		t.infants = float32(b.Infants) * pricing["heliride"].infants
		return t
	}

	switch b.Package {
	case "jatayu":
		t = packageJatayu()
	case "adventure":
		t = packageAdventure()
	case "heliride":
		t = packageHeliride()
	}

	b.Amount = t.adults + t.children + t.infants
	fmt.Println(b)
	tx := c.Value("tx").(*pop.Connection)
	err := tx.Save(&b)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(err.Error())

	return c.Redirect(302, "/")

	return c.Render(200, r.HTML("booking/book.html"))
}
