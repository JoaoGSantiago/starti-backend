package handlers

import (
	"net/http"

	"github.com/JoaoGSantiago/starti-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var input services.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.service.Create(input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusCreated, post)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	post, err := h.service.GetByID(id)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	successResponse(c, http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var input services.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.service.Update(id, input)
	if err != nil {
		if err.Error() == "post nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusOK, post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	if err := h.service.Delete(id); err != nil {
		if err.Error() == "post nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *PostHandler) ArchivePost(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	if err := h.service.Archive(id); err != nil {
		if err.Error() == "post nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusOK, gin.H{"message": "post archived"})
}

func (h *PostHandler) ListPostComments(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	comments, err := h.service.ListComments(id)
	if err != nil {
		if err.Error() == "post nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusOK, comments)
}
