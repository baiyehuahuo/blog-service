package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

func init() {
	var err error
	if err = setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err = setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	if err = setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目 第二章
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	global.Logger.Infof("%s: test-logger/%s", "fwf", "blog-service")
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(global.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func setupSetting() error {
	st, err := setting.NewSetting()
	if err != nil {
		return err
	}

	if err = st.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	log.Printf("%#v", global.ServerSetting)

	if err = st.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	log.Printf("%#v", global.AppSetting)

	if err = st.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	log.Printf("%#v", global.DatabaseSetting)

	if err = st.ReadSection("JWT", &global.JWTSetting); err != nil {
		return err
	}
	log.Printf("%#v", global.JWTSetting)

	if err = st.ReadSection("Email", &global.EmailSetting); err != nil {
		return err
	}
	log.Printf("%#v", global.EmailSetting)

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
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
		Filename:  path.Join(global.AppSetting.LogSavePath, global.AppSetting.LogFileName+global.AppSetting.LogFileExt),
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
