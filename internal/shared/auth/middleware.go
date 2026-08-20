package middleware

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kadr/globe_express/config"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func RequestID(c fiber.Ctx) error {
	if req := c.Get("X-Request-Id"); req == "" {
		c.Set("X-Request-Id", uuid.NewString())
	}
	return c.Next()
}

func Logger(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		log := logger.With(
			slog.String("request_id", c.GetRespHeader("X-Request-Id")),
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
		)
		c.Locals("logger", log)
		return c.Next()
	}
}

func Recovery(c fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			if logger, ok := c.Value("logger").(*slog.Logger); ok {
				logger.Error("panic recovery", r)
			} else {
				slog.Error("panic recovery", r)
			}
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": r})
		}
	}()
	return c.Next()
}

func GetUserFromReq(appConfig *config.AppConfig) fiber.Handler {
	return func(c fiber.Ctx) error {
		var logger *slog.Logger
		if log, ok := c.Value("logger").(*slog.Logger); ok {
			logger = log
		} else {
			logger = slog.Default()
		}
		jwtAuth := NewJWT(appConfig)
		if token := c.Get("Authorization"); token != "" {
			if err := jwtAuth.checkValid(token); err != nil {
				logger.Error(err.Error())
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
			}
			userID, err := jwtAuth.getUserID(token)
			if err != nil {
				logger.Error(err.Error())
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
			}
			c.Locals("userID", userID)
			if userRole := jwtAuth.getUserRole(token); userRole != "" {
				c.Locals("userRole", userRole)
			}
		} else {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
		}

		return c.Next()
	}
}

func ExecuteResponse(c fiber.Ctx) error {
	err := c.Next()
	logger := c.Value("logger").(*slog.Logger)
	if err != nil {
		switch {
		case errors.Is(err, api_errors.ErrorNotFound):
			logger.Warn("Not found record")
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		case errors.Is(err, api_errors.ErrorBadRequest):
			logger.Warn(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, api_errors.ErrorIncorrectStatus):
			logger.Warn(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, api_errors.ErrorCanNotCancel):
			logger.Warn(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, api_errors.ErrorFieldValidation):
			logger.Warn(err.Error())
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, api_errors.ErrorForbidden):
			logger.Warn(err.Error())
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		default:
			logger.Warn(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
		}
	}
	return nil
}
