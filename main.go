package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"web_app/api"
	_ "web_app/docs"
	"web_app/middleware"
	"web_app/settings"
	"web_app/tool"
)

// @title web_app项目接口文档
// @version 1.0
// @description 投票帖子网站后端接口
// @contact.name cold bin
// @contact.url https://github.com/liuhaibin123456789/web-app.git
// @host 127.0.0.1:8085
// @securityDefinitions.apikey CoreAPI
// @name Authorization
// @in header
// @BasePath /api/v1
func main() {
	settings.Settings()
	URL()
}

func URL() {

	//注册路由
	router := gin.New()

	router.Use(middleware.GinLogger, middleware.GinZapRecovery(true), middleware.RateLimitMiddleware(200, 1))
	router.Use(middleware.Cors())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //注册gin-swagger路由

	urls(router)
	err := router.RunTLS(":"+tool.V.GetString("app.port"), "./config/cert.pem", "./config/key.pem")
	if err != nil {
		log.Println(err)
		return
	}
}

func urls(r *gin.Engine) {
	//加载静态资源
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := r.Group("/api/v1")
	v1.POST("register", api.Register)
	v1.POST("login", api.Login)
	v1.POST("tokens", api.GetTokens) //本路由只能在需要刷新双token时使用

	v1.Use(middleware.Jwt())
	{
		v1.GET("community", api.GetCommunity)
		v1.POST("community", api.CreateCommunity)
		v1.GET("community/:community_id", api.GetDetailCommunity)
		v1.POST("post", api.CreatePost)
		//v1.GET("post", api.GetPost)   //分页查询
		v1.GET("post2", api.GetPost2) //分页查询
		v1.GET("post/:post_id", api.GetDetailPost)
		v1.POST("vote", api.VoteForPost)
	}
}
