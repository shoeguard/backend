package utils

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidateForm(c *fiber.Ctx, form interface{}) error {
	if err := c.BodyParser(&form); err != nil {
		return err
	}

	// validating body
	if succeed, err := govalidator.ValidateStruct(form); !succeed {
		return err
	}

	return nil
}
