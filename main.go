package main

import (
	"log"
	"login-demo/config"
	"login-demo/database"
	"login-demo/handlers"
	"login-demo/middleware"
	"login-demo/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	if err := database.InitDB(cfg); err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移表结构
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 初始化 Gin
	r := gin.Default()

	// 初始化处理器
	authHandler := handlers.NewAuthHandler(database.DB, cfg.JWTSecret)

	// 路由设置
	api := r.Group("/api")
	{
		// 公开路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// 需要认证的路由
		protected := api.Group("/user")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			protected.GET("/profile", authHandler.GetProfile)
		}
	}

	// 启动服务器
	log.Println("服务器启动在 :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
