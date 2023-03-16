package routes

import (
	"test-golang/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// route register login
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// route notes
	app.Get("/api/notes", controllers.GetNotes)
	app.Get("/api/note/:NoteID", controllers.GetNoteById)
	app.Put("/api/update/:NoteID", controllers.UpdateNoteById)
	app.Post("/api/create", controllers.CreateNote)
	app.Delete("/api/delete/:NoteID", controllers.DeleteNote)

}
