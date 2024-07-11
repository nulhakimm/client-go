package controller

import (
	"github.com/faridlan/client-go/model"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	UserRender(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	UserModel model.UserModel
}

func NewUserController(userModel model.UserModel) UserController {
	return &UserControllerImpl{
		UserModel: userModel,
	}
}

func (controller *UserControllerImpl) UserRender(ctx *fiber.Ctx) error {

	userResponse, err := controller.UserModel.FindAll(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.Render("index", userResponse)

}
