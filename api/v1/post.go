package v1

import (
	"net/http"
	"strconv"

	"github.com/TemurMannonov/blog/api/models"
	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /posts/{id} [get]
// @Summary Get post by id
// @Description Get post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.storage.Post().Get(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	post := parsePostModel(resp)

	likesInfo, err := h.storage.Like().GetLikesDislikesCount(post.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post.LikeInfo = &models.PostLikeInfo{
		LikesCount:    likesInfo.LikesCount,
		DislikesCount: likesInfo.DislikesCount,
	}

	c.JSON(http.StatusOK, post)
}

// @Security ApiKeyAuth
// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		req models.CreatePostRequest
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

	resp, err := h.storage.Post().Create(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      payload.UserID,
		CategoryID:  req.CategoryID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	post := parsePostModel(resp)
	c.JSON(http.StatusCreated, post)
}

// @Router /posts [get]
// @Summary Get all posts
// @Description Get all posts
// @Tags post
// @Accept json
// @Produce json
// @Param filter query models.GetAllPostsParams false "Filter"
// @Success 200 {object} models.GetAllPostsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	req, err := validateGetAllPostsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Post().GetAll(&repo.GetAllPostsParams{
		Page:       req.Page,
		Limit:      req.Limit,
		Search:     req.Search,
		UserID:     req.UserID,
		CategoryID: req.CategoryID,
		SortByData: req.SortByData,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, getPostsResponse(result))
}

func validateGetAllPostsParams(c *gin.Context) (*models.GetAllPostsParams, error) {
	var (
		limit              int = 10
		page               int = 1
		err                error
		userID, categoryID int
		sortByDate         string = "desc"
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("user_id") != "" {
		userID, err = strconv.Atoi(c.Query("user_id"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("category_id") != "" {
		categoryID, err = strconv.Atoi(c.Query("category_id"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("sort_by_date") != "" &&
		(c.Query("sort_by_date") == "desc" || c.Query("sort_by_date") == "asc") {
		sortByDate = c.Query("sort_by_date")
	}

	return &models.GetAllPostsParams{
		Limit:      int32(limit),
		Page:       int32(page),
		Search:     c.Query("search"),
		UserID:     int64(userID),
		CategoryID: int64(categoryID),
		SortByData: sortByDate,
	}, nil
}

func getPostsResponse(data *repo.GetAllPostsResult) *models.GetAllPostsResponse {
	response := models.GetAllPostsResponse{
		Posts: make([]*models.Post, 0),
		Count: data.Count,
	}

	for _, post := range data.Posts {
		p := parsePostModel(post)
		response.Posts = append(response.Posts, &p)
	}

	return &response
}

func parsePostModel(post *repo.Post) models.Post {
	return models.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserID:      post.UserID,
		CategoryID:  post.CategoryID,
		CreatedAt:   post.CreatedAt,
		ViewsCount:  post.ViewsCount,
	}
}
