package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/controllers"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/user", controllers.User)
	api.Post("/logout", controllers.Logout)
	api.Patch("/user/update/:id", controllers.UpdateUser)

	api.Get("/pengaduan", controllers.IndexPengaduan)
	api.Post("/pengaduan/create", controllers.CreatePengaduan)
	api.Get("/pengaduan/:id", controllers.ShowPengaduan)
	api.Patch("/pengaduan/update/:id", controllers.UpdatePengaduan)
	api.Delete("/pengaduan/delete/:id", controllers.DeletePengaduan)

	api.Post("/tanggapan/create/:id", controllers.CreateTanggapan)

	api.Post("/image/upload", controllers.UploadImageTest)
}
