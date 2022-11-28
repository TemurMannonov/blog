package v1

import (
	"net/http"
	"strconv"

	"github.com/TemurMannonov/blog/api/models"
	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /likes [post]
// @Summary Create or update like
// @Description Create or update like
// @Tags like
// @Accept json
// @Produce json
// @Param like body models.CreateOrUpdateLikeRequest true "like"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateOrUpdateLike(c *gin.Context) {
	var (
		req models.CreateOrUpdateLikeRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = h.storage.Like().CreateOrUpdate(&repo.Like{
		UserID: payload.UserID,
		PostID: req.PostID,
		Status: req.Status,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully finished",
	})
}

// @Security ApiKeyAuth
// @Router /likes/user-post [get]
// @Summary Get like by user and post
// @Description Get like by user and post
// @Tags like
// @Accept json
// @Produce json
// @Param post_id query int true "Post ID"
// @Success 200 {object} models.Like
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetLike(c *gin.Context) {
	postID, err := strconv.Atoi(c.Query("post_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp, err := h.storage.Like().Get(payload.UserID, int64(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.Like{
		ID:     resp.ID,
		PostID: resp.PostID,
		UserID: resp.UserID,
		Status: resp.Status,
	})
}
