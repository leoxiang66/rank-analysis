package main

import (
	"context"
	"fmt"
	"lol_professor/backend/automation"
	"time"
	"lol_professor/backend/api/handlers"
	"github.com/gin-gonic/gin"
	"lol_professor/backend/api"
)


var Gin_ctx *gin.Context

// App struct
type App struct {
	ctx       context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{

	}
}

// reverseSlice reverses a slice of any type T
func reverseSlice[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i] // Swap elements
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	r := gin.Default()
	r.Use(handlers.Cors())

	// 初始化路由
	api.InitRoutes(r)
	go automation.StartAutomation()
	// 启动服务
	r.Run(":11451") // 在 11451 端口上运行
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Info(name string) {
	fmt.Println(name)
}


func (a *App) waitOnClient(){
	for {
		time.Sleep(time.Second)
		// handlers.GetSummonerAndRank()
	}
}