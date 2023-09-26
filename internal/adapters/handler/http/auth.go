package AuthHandler

import (
	"deluze/internal/core/ports"
	"deluze/pkg/common"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type httpAuthHandler struct {
	service ports.AuthService
}

func NewHttpAuthHandler(service ports.AuthService) *httpAuthHandler {
	return &httpAuthHandler{
		service: service,
	}
}

type LoginBody struct {
	Email string `json:"email"`
}

func (h *httpAuthHandler) Login(c *fiber.Ctx) error {
	body := new(LoginBody)
	c.BodyParser(body)
	err := h.service.Login(body.Email)
	if err != nil {
		return common.HttpErrorResponse(c, 500, err.Error())
	}
	return c.JSON(fiber.Map{
		"message": "an verification code send to users email",
	})
}

type VerifyQuery struct {
	Email string `query:"email"`
	Code  string `query:"code"`
}

func (h *httpAuthHandler) Verify(c *fiber.Ctx) error {
	query := new(VerifyQuery)
	c.QueryParser(query)
	user, token, err := h.service.Verify(query.Email, query.Code)
	if err != nil {
		return common.HttpErrorResponse(c, 500, err.Error())
	}
	c.Cookie(&fiber.Cookie{
		Name:     "_auth",
		Value:    *token,
		Path:     "/",
		Domain:   os.Getenv("APP_HOST"),
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	})
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "login successfully",
	})
}

func (h *httpAuthHandler) Logout(c *fiber.Ctx) error {
	c.ClearCookie("_auth")
	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}
