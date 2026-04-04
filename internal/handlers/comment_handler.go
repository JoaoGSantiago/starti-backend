package handlers

import (
	"net/http"

	"github.com/JoaoGSantiago/starti-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service services.CommentService
}

func NewCommentHandler(service services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var input services.CreateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.service.Create(input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusCreated, comment)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var input services.UpdateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comment, err := h.service.Update(id, input)
	if err != nil {
		if err.Error() == "comentario nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(c, http.StatusOK, comment)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	if err := h.service.Delete(id); err != nil {
		if err.Error() == "comentario nao encontrado" {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
