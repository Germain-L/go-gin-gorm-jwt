package cmd

import (
	"fmt"
	"server/database"
	"server/handlers"
	"server/middlewares"
	"server/utils"

	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, log bool) *gin.Engine {

	r := gin.Default()

	m := middlewares.NewMiddleware(db)
	h := handlers.NewHandler(db)

	if log {
		r.Use(m.Logger())
	}

	r.GET("/env", h.Env)

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
	}

	api := r.Group("/api")
	api.Use(m.Authorize())
	{
		api.GET("/me", h.GetUser)
	}

	return r
}

func Run() {
	utils.LoadEnv()

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	dialector := postgres.Open(dsn)
	db, err := database.ConnectToDB(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := SetupRouter(db, true)
	r.Run(":8080")
}
