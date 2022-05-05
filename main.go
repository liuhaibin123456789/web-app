package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/api"
	"web_app/dao/mysql"
	_ "web_app/docs"
	"web_app/middleware"
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
	//加载配置
	if err := tool.Viper(); err != nil {
		fmt.Println("viper出错:", err)
		return
	}
	//初始化日志
	if err := tool.Logger(); err != nil {
		fmt.Println("zap出错:", err)
		return
	}
	tool.SugaredDebug("zap logger初始化...")

	if err := mysql.Mysql(); err != nil {
		tool.SugaredPanicf("mysql init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("mysql 初始化...")

	if err := tool.Redis(); err != nil {
		tool.SugaredPanicf("redis init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("redis 初始化...")

	if err := tool.InitSnowflake(); err != nil {
		tool.SugaredPanicf("snowflake init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("snowflake 节点初始化...")
	//路由
	URL()
}

func URL() {

	//注册路由
	router := gin.New()

	router.Use(middleware.GinLogger, middleware.GinZapRecovery(true))
	router.Use(middleware.Cors())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //注册gin-swagger路由

	urls(router)

	//优雅关机及重启
	srv := &http.Server{Addr: "localhost:" + tool.V.GetString("app.port"), Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			tool.SugaredPanicf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)                      // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	tool.SugaredDebug("Shutdown Server ...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		tool.SugaredFatal("Server Shutdown: ", err)
	}

}

func urls(r *gin.Engine) {

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
