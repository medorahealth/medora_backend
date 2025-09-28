package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/medorahealth/medora_backend/internal/model"
	"github.com/medorahealth/medora_backend/internal/service"
)

// LabHandler holds the dependencies for the lab HTTP handlers.
type LabHandler struct {
	labService service.LabService
	validate   *validator.Validate
}

// NewLabHandler creates a new lab handler.
func NewLabHandler(labService service.LabService) *LabHandler {
	return &LabHandler{
		labService: labService,
		validate:   validator.New(),
	}
}

// RegisterLab is the HTTP handler for creating a new lab.
func (h *LabHandler) RegisterLab(c *gin.Context) {
	var req model.RegisterLabRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed: " + err.Error()})
		return
	}

	lab, err := h.labService.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lab"})
		return
	}

	c.JSON(http.StatusCreated, lab)
}
