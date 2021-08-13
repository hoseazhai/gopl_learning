package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hoseazhai/gopl_learning/blog-service/global"
	"github.com/hoseazhai/gopl_learning/blog-service/internal/model"
	"github.com/hoseazhai/gopl_learning/blog-service/internal/routers"
	"github.com/hoseazhai/gopl_learning/blog-service/pkg/logger"
	"github.com/hoseazhai/gopl_learning/blog-service/pkg/setting"
	"github.com/hoseazhai/gopl_learning/blog-service/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "192.168.20.129:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil

	return nil
}

// @title 博客测试系统
// @version 1.1
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/gopl_learning/blog-server
func main() {
	var test context.Context
	global.Logger.Infof(test, "eddycjy", "blog-service")
	gin.SetMode(global.ServerSetting.RunMode)
	route := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           route,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout:      global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}
