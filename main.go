package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"mch2022/config"
	"mch2022/db"
	"mch2022/handlers/nko"
	"mch2022/handlers/user"
	"net/http"
)

func main() {

	config.Load()
	db.Connect()

	app := gin.Default()
	app.Use(cors.Default())

	gin.SetMode(gin.DebugMode)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	app.GET("", func(c *gin.Context) {
		if db.Ping() == nil {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "db error"})
		return
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	usersGroup := v1.Group("/users")
	nkoGroup := v1.Group("/nko")
	projectsGroup := v1.Group("/projects")

	nkoGroup.Use(user.Middleware)
	usersGroup.Use(user.Middleware)
	projectsGroup.Use(user.Middleware)

	usersGroup.POST("/register", user.Register)
	usersGroup.POST("/auth", user.Auth)
	nkoGroup.POST("/auth", user.AuthNKO)
	nkoGroup.POST("/register", nko.RegisterNKO)

	projectsGroup.GET("", nko.FindProjects)
	projectsGroup.GET("/:projectID", nko.GetProjectByID)
	projectsGroup.GET("/search", nko.SearchProjects)

	nkoGroup.GET("", nko.FindNkos)
	nkoGroup.GET("/:id", nko.GetNKObyID)
	usersGroup.GET("/:id", user.GetUserById)
	usersGroup.GET("/global_rating", user.GetWholeRatingList)

	usersGroup.GET("", user.GetUser)
	nkoGroup.GET("/me", user.GetNKO)
	usersGroup.PATCH("/project/:id", user.ProjectParticipate)
	usersGroup.DELETE("/project/:id", user.DeleteParticipant)

	nkoGroup.GET("/projects", user.GetNKOProjects)

	nkoGroup.POST("/project", nko.CreateProject)

	nkoGroup.POST("/:userID/rate", nko.RateUser)
	usersGroup.POST("/:nkoID/rate", user.RateNKO)

	projectsGroup.GET("/:projectID/participants", user.ProjectParticipants) // admin only - get event users
	projectsGroup.GET("/filters", nko.GetFiltersRefs)

	err := app.Run("0.0.0.0:" + config.ServerPort)
	if err != nil {
		log.Println(err)
	}
}
