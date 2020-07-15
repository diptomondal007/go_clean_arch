package http

import (
	"github.com/diptomondal007/go_clean_arch/auth"
	"github.com/diptomondal007/go_clean_arch/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Handler struct {
	usecase auth.UseCase
}

type authInput struct {
	Username string
	Password string
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		usecase: useCase,
	}
}

func (h *Handler) SignIn(c echo.Context) error {
	inp := new(authInput)
	if err := c.Bind(&inp); err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Message: "inputs are not valid",
		})
	}

	token, err := h.usecase.SignIn(c.Request().Context(), inp.Username, inp.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			return c.JSON(http.StatusUnauthorized, &Response{
				Message: err.Error(),
			})
		}
		if err == auth.ErrPasswordDoesntMatch {
			return c.JSON(http.StatusUnauthorized, &Response{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "log in successful",
		Data:    token,
	})
}

func (h *Handler) SignUp(c echo.Context) error {
	inp := new(authInput)
	if err := c.Bind(inp); err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "sign up failed",
		})
	}

	var user *models.User
	user = &models.User{
		Username:  inp.Username,
		Password:  inp.Password,
		CreatedAt: time.Now().Format(time.RFC850),
	}

	if err := h.usecase.SignUp(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "sign up failed",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "sign up successful",
	})
}
