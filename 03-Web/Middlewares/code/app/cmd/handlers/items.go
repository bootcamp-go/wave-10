package handlers

import (
	"app/internal/domain"
	"app/internal/item"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ItemJSON is an struct that represents an item in JSON format.
type ItemJSON struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// NewItemsDefault returns a new instance of ItemsDefault.
func NewItemsDefault(sv item.Service) *ItemsDefault {
	return &ItemsDefault{
		sv: sv,
	}
}

// ItemsDefault is an struct that contains handlers to manage items.
type ItemsDefault struct {
	// sv is the service used by the handler.
	sv item.Service
}

// ShowGetItems godoc
// @Summary Show all items
// @Description get all items
// @Tags items
// @Produce json
// @Failure 500 {object} map[string]any
// @Success 200 {object} map[string]any
// @Router /items [get]
func (h *ItemsDefault) GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		// ...

		// process
		// - get all items
		items, err := h.sv.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}

		// response
		data := make([]ItemJSON, len(items))
		for i, item := range items {
			data[i] = ItemJSON{
				ID:    item.ID,
				Name:  item.Name,
				Price: item.Price,
			}
		}
		c.JSON(http.StatusOK, map[string]any{"message": "success to find items", "data": data})
	}
}

// GetItem returns the item with the given ID.
// @Summary Get an item
// @Description get an item
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Failure 400 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Success 200 {object} map[string]any
// @Router /items/{id} [get]
func (h *ItemsDefault) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		// - get the ID from the URL
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"message": "invalid ID"})
			return
		}

		// process
		// - get the item
		it, err := h.sv.FindByID(id)
		if err != nil {
			switch {
			case errors.Is(err, item.ErrServiceItemNotFound):
				c.JSON(http.StatusNotFound, map[string]any{"message": "item not found"})
			default:
				c.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		data := ItemJSON{
			ID:    it.ID,
			Name:  it.Name,
			Price: it.Price,
		}
		c.JSON(http.StatusOK, map[string]any{"message": "success to find item", "data": data})
	}
}

// BodyRequest is an struct that represents the request body to create an item.
type BodyRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// PostItem saves the given item.
// @Summary Save an item
// @Description save an item
// @Tags items
// @Accept json
// @Produce json
// @Param body body BodyRequest true "Item Body"
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Success 200 {object} map[string]any
// @Router /items [post]
func (h *ItemsDefault) PostItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		var body BodyRequest
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}

		// process
		// - serialize item
		it := domain.Item{
			Name:  body.Name,
			Price: body.Price,
		}
		// - save the item
		err = h.sv.Save(&it)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}

		// response
		// - serialize to JSON
		data := ItemJSON{
			ID:    it.ID,
			Name:  it.Name,
			Price: it.Price,
		}
		c.JSON(http.StatusOK, map[string]any{"message": "success to save item", "data": data})
	}
}