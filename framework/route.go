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

	r := gin.Default()
	// Khởi tạo handler
	loginHandler := handler.NewLoginHandler()

	// Route đăng nhập không cần xác thực
	r.GET("/login", loginHandler.Login)

	// Nhóm các route cần xác thực
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		meterHandler := handler.NewMeterHandler()
		// Route yêu cầu đã đăng nhập
		authorized.GET("/data", meterHandler.GetMeter)
		meterTodayHandler := handler.NewMeterHandlerToday()
		authorized.GET("/datatoday", meterTodayHandler.GetMeterToday)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
