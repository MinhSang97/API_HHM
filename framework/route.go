package framework

import (
	"app/dbutil"
	"app/handler"
	"app/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Route() {
	db, err := dbutil.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	fmt.Println("Connected: ", db)

	// CRUD: Create, Read, Update, Delete
	// POST /v1/items (create a new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT | PATCH) /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (delete item by id)
	// viper.SetConfigFile("config.yaml")
	// if err := viper.ReadInConfig(); err != nil {
	// 	panic(err)
	// }

	// Your existing code to set up routes and database

	// Register your route, passing the configuration

	r := gin.Default()
	// Khởi tạo handler
	loginHandler := handler.NewLoginHandler()

	// Route đăng nhập không cần xác thực
	r.GET("/login", loginHandler.Login)

	// Nhóm các route cần xác thực
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Route yêu cầu đã đăng nhập
		authorized.GET("/data", loginHandler.GetMeter)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
