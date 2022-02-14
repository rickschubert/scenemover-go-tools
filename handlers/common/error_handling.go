package common

import "github.com/gofiber/fiber/v2"

type Error struct {
	Message string `json:"message"`
}

func SendError(err error, context *fiber.Ctx) error {
	return context.Status(fiber.ErrBadGateway.Code).JSON(Error{
		Message: err.Error(),
	})
}
