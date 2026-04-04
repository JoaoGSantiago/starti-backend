package handlers

import (
	"net/http"
	"strconv"

	"github.com/JoaoGSantiago/starti-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListAll()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	successResponse(c, http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input services.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Create(input)
	if err != nil {
		if err.Error() == "email em uso" || err.Error() == "username em uso" {
			errorResponse(c, http.StatusConflict, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	successResponse(c, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var input services.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.Update(id, input)
	if err != nil {
		if err.Error() == "usuario nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	if err := h.service.Delete(id); err != nil {
		if err.Error() == "usuario nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) ListUserPosts(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	posts, err := h.service.ListPublicPosts(id)
	if err != nil {
		if err.Error() == "usuario nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusOK, posts)
}

func (h *UserHandler) ListUserComments(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	comments, err := h.service.ListPublicComments(id)
	if err != nil {
		if err.Error() == "usuario nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusOK, comments)
}

func parseID(c *gin.Context, param string) (uint, error) {
	raw := c.Param(param)
	id, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "id invalido")
		return 0, err
	}
	return uint(id), nil
}
