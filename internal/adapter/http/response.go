package http

import (
	"fiber-server-1/internal/core/models"

	"github.com/gofiber/fiber/v2"
)

var errorStatusMap = map[error]int{
	models.ErrDataNotFound:   fiber.ErrNotFound.Code,
	models.ErrInternalServer: fiber.ErrInternalServerError.Code,
}
