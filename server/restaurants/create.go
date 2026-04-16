package restaurants

import (
	"chooseMyRestaurant/models"
	"chooseMyRestaurant/repo/dto"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRepo interface {
	CreateRestaurant(ctx context.Context, request dto.CreateRestaurantInput) (id string, err error)
	FetchRestaurant(ctx context.Context, id string) (dto.Restaurant, error)
}

type createHandler struct {
	repo CreateRepo
}

type createRequest struct {
	models.Restaurant
}

func (h *createHandler) Handle(c *gin.Context) {
	var (
		req = c.MustGet("request").(*createRequest)
	)
	var input dto.CreateRestaurantInput
	input.FromModel(req.CreateRestaurantInput)
	id, err := h.repo.CreateRestaurant(c, input)
	if err != nil {
		if errors.Is(err, repo.ErrConflict) {
			c.JSON(http.StatusConflict, gin.H{
				"error", "restaurant already exists",
			})
			return
		}
	}

	r, err := h.repo.FetchRestaurant(c, id)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Could not find restaurant",
			})
			return
		}
	}

}

func (h *createHandler) Validate(c *gin.Context) {
	var req createRequest
	if !req.bindAndValidate(c) {
		return
	}
	c.Set("request", &req)

}

func (r *createRequest) bindAndValidate(c *gin.Context) bool {
	if err := c.BindJSON(r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to bind body as JSON: %v", err),
		})
		return false
	}
	if r.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request should have a name",
		})
		return false
	}
	return true
}
