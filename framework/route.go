package framework

import (
	"app/dbutil"
	"app/handler"
	"app/middleware"
	"context"
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
	r.Use(middleware.ErrorHandler())

	// r.Use(middleware.BasicAuthMiddleware())

	r.GET("/login", func(c *gin.Context) {
		API_User := c.Query("user")
		API_PassWord := c.Query("password")

		r.Use(middleware.ValidLogin(context.Background(), db, API_User, API_PassWord))

		handler.Login(db)(c)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
