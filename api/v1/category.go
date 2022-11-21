package v1

import (
	"net/http"
	"strconv"

	"github.com/TemurMannonov/blog/api/models"
	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /categories/{id} [get]
// @Summary Get category by id
// @Description Get category by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Category().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.Category{
		ID:        resp.ID,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		req models.CreateCategoryRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp, err := h.storage.Category().Create(&repo.Category{Title: req.Title})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.Category{
		ID:        resp.ID,
		Title:     resp.Title,
		CreatedAt: resp.CreatedAt,
	})
}

// @Router /categories [get]
// @Summary Get all categories
// @Description Get all categories
// @Tags category
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 200 {object} models.GetAllCategoriesResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllCategories(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Category().GetAll(&repo.GetAllCategoriesParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getCategoriesResponse(result))
}

func getCategoriesResponse(data *repo.GetAllCategoriesResult) *models.GetAllCategoriesResponse {
	response := models.GetAllCategoriesResponse{
		Categories: make([]*models.Category, 0),
		Count:      data.Count,
	}

	for _, c := range data.Categories {
		response.Categories = append(response.Categories, &models.Category{
			ID:        c.ID,
			Title:     c.Title,
			CreatedAt: c.CreatedAt,
		})
	}

	return &response
}
