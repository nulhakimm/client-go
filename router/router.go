package router

import (
	"github.com/faridlan/client-go/controller"
	"github.com/faridlan/client-go/model"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {

	userModel := model.NewUserModel()
	userController := controller.NewUserController(userModel)

	app.Get("/", userController.UserRender)
}
