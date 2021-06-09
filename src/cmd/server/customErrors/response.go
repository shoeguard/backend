package customErrors

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Error  CustomError
	Detail string
}

func Response400WithError(
	c *fiber.Ctx,
	customError CustomError,
	detail string,
) error {
	return c.Status(400).JSON(ErrorResponse{customError, detail})
}
