package adapters

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"products/internal/application/ports"
)

type ProductHandler struct {
	CreateUC ports.ProductCreationInputPort
}

func NewProductHandler(uc ports.ProductCreationInputPort) *ProductHandler {
	return &ProductHandler{CreateUC: uc}
}

type createProductRequest struct {
	Name        string  `json:"name"        binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price"       binding:"required"`
}

type productResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// POST /products
func (h *ProductHandler) Create(c *gin.Context) {
	var req createProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body: " + err.Error()})
		return
	}

	cmd := ports.ProductCreationCommand{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	p, err := h.CreateUC.Create(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	resp := productResponse{
		ID:          p.ID.String(),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
	}
	c.JSON(http.StatusCreated, resp)
}

func NewRouter(h *ProductHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.POST("/products", h.Create)
	return r
}
