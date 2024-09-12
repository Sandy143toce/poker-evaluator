package utils

import (
	"net/url"

	"github.com/Sandy143toce/poker-evaluator/backend/models"
	"github.com/gofiber/fiber/v2"
)

func AuthorizeRequest(structName string) fiber.Handler {
	// Authorize the request
	return func(c *fiber.Ctx) error {
		errBody, valid := RequestValidation(c, structName)
		if !valid {
			return c.Status(fiber.StatusBadRequest).JSON(errBody)
		}
		return c.Next()
	}
}

func RequestValidation(c *fiber.Ctx, structName string) (models.ErrorResponse, bool) {

	switch structName {
	case "CreateWebhookRequest":

	case "AddCustomerEndpointRequest":

	case "SendEventRequest":
	}
	return models.ErrorResponse{}, true
}

func isValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
