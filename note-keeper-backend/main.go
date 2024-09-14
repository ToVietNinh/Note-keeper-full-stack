package main

import (
	"log"
	"os"

	"aedronef1_backend/config"
	"aedronef1_backend/docs"
	"aedronef1_backend/infrastucture/repository"

	"aedronef1_backend/usecase/user"

	authHandler "aedronef1_backend/api/handler/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.SetConfigFile("config")
}

func main() {
	// envConfig := getConfig()
	err := gotenv.Load()
	if err != nil {
		log.Println(err)
	}

	// Database
	db, err := repository.ConnectDatabase(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	if err != nil {
		log.Println(err)
		return
	}

	// // Redis
	// redisClient, err := repository.ConnectRedis(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// App
	app := gin.New()
	// Cors
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})
	app.Use(crs)

	if err != nil {
		log.Println(err)
		return
	}

	// Define Repository
	userRepo := repository.NewUserRepository(db)
	statusRepo := repository.NewStatusRepository(db)

	// Define Service
	userService := user.NewService(userRepo, statusRepo)

	// Handler
	authHandler.MakeHandlers(app, userService)

	docs.SwaggerInfo.BasePath = ""
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// app.Run(fmt.Sprintf("%s%s%v", os.Getenv("APP_HOST"), ":", os.Getenv("APP_PORT")))
	app.Run()

}
