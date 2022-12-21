package api

import (
	v1 "github.com/TemurMannonov/blog/api/v1"
	"github.com/TemurMannonov/blog/config"
	"github.com/TemurMannonov/blog/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/TemurMannonov/blog/api/docs" // for swagger
)

type RouterOptions struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
}

// @title           Swagger for blog api
// @version         1.0
// @description     This is a blog service api.
// @BasePath  /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:      opt.Cfg,
		Storage:  opt.Storage,
		InMemory: opt.InMemory,
	})

	router.Static("/media", "./media")

	apiV1 := router.Group("/v1")

	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.POST("/users", handlerV1.AuthMiddleware, handlerV1.CreateUser)
	apiV1.GET("/users", handlerV1.GetAllUsers)
	apiV1.GET("/users/me", handlerV1.AuthMiddleware, handlerV1.GetUserProfile)

	apiV1.GET("/categories/:id", handlerV1.GetCategory)
	apiV1.POST("/categories", handlerV1.AuthMiddleware, handlerV1.CreateCategory)
	apiV1.GET("/categories", handlerV1.GetAllCategories)
	apiV1.PUT("/categories/:id", handlerV1.AuthMiddleware, handlerV1.UpdateCategory)
	apiV1.DELETE("/categories/:id", handlerV1.AuthMiddleware, handlerV1.DeleteCategory)

	apiV1.GET("/posts/:id", handlerV1.GetPost)
	apiV1.POST("/posts", handlerV1.AuthMiddleware, handlerV1.CreatePost)
	apiV1.GET("/posts", handlerV1.GetAllPosts)

	apiV1.POST("/comments", handlerV1.AuthMiddleware, handlerV1.CreateComment)
	apiV1.GET("/comments", handlerV1.GetAllComments)

	apiV1.POST("/likes", handlerV1.AuthMiddleware, handlerV1.CreateOrUpdateLike)
	apiV1.GET("/likes/user-post", handlerV1.AuthMiddleware, handlerV1.GetLike)

	apiV1.POST("/auth/register", handlerV1.Register)
	apiV1.POST("/auth/verify", handlerV1.Verify)
	apiV1.POST("/auth/login", handlerV1.Login)
	apiV1.POST("/auth/forgot-password", handlerV1.ForgotPassword)
	apiV1.POST("/auth/verify-forgot-password", handlerV1.VerifyForgotPassword)
	apiV1.POST("/auth/update-password", handlerV1.AuthMiddleware, handlerV1.UpdatePassword)

	apiV1.POST("/file-upload", handlerV1.AuthMiddleware, handlerV1.UploadFile)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
