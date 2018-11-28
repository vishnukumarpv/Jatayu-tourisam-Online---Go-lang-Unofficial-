package actions

import (
	"database/sql"
	"projects/gojatayu/models"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func RegHandler(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		errors.WithStack(err)
	}

	// u.CreatedAt = time.Now().Local()
	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		// return c.Render(422, r.HTML("users/register.html"))
		c.Flash().Add("error", "Account Error occured!!!.")
		return c.Redirect(302, "/")
	}
	c.Flash().Add("success", "Account created successfully.")
	c.Session().Set("user_id", u.ID)
	c.Flash().Add("success", "Welcome to jatayu tourisam.")
	// and redirect to the home page
	return c.Redirect(302, "/")
}

func LoginHandler(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	err := tx.Where("email = ?", strings.ToLower(u.Email)).First(u)

	bad := func() error {
		c.Set("user", u)
		verrs := validate.NewErrors()
		verrs.Add("email", "invalid email/password")
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("auth/new.html"))
	}
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return bad()
		}
		return errors.WithStack(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return bad()
	}
	c.Session().Set("user_id", u.ID)
	c.Flash().Add("success", "Welcome Back to Jatayu tourisam!")

	return c.Redirect(302, "/")

	// return c.Render(200, r.JSON(u))
}

func LogoutHandler(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out!")
	return c.Redirect(302, "/")
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
