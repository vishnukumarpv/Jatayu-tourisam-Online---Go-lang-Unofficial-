package actions

import "github.com/gobuffalo/buffalo"

func routes() *buffalo.App {
	// app.GET("/booking", BookingHandler)
	app.GET("/profile", ProfileHandler)
	app.GET("/ticket", TicketHandler)
	// app.GET("/login", LoginHandler)

	app.POST("/login", LoginHandler)
	app.GET("/logout", LogoutHandler)
	app.POST("/register", RegHandler)

	auth := app.Group("/pi")
	book := app.Group("/booking")

	book.GET("/", BookingHandler)
	book.POST("/book", BookHandler)

	auth.Use(Authorize)
	// auth.GET("/booking/book", BookingBook)

	return app
}
