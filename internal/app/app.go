package app

import (
	"context"
	"go-blog-web/internal/config"
	"go-blog-web/internal/handler"
	"go-blog-web/internal/middleware"
	"go-blog-web/internal/router"
	"go-blog-web/internal/service"
	"go-blog-web/internal/storage/memory"
	"go-blog-web/internal/util"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln("Ошибка загрузки настроек", err)
		return
	}
	log := util.InitLogger(cfg.Env)

	storage := memory.NewStorage()
	service := service.NewService(storage)
	handler := handler.NewHandler(service, log)

	if err := service.CreateFirstAdmin(cfg); err != nil {
		log.Error("Не удалось создать администратора", util.Err(err))
		return
	}

	server := &http.Server{
		Addr:    cfg.ServerAddress(),
		Handler: middleware.CORSMiddleware(cfg.Cors(), router.NewRouter(handler, service)),
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Запуск сервера", slog.String("address", cfg.ServerAddress()))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Ошибка запуска сервера", util.Err(err))
		}
	}()

	<-sigChan
	log.Info("Получен сигнал завершения, начинаем плавное завершение...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Ошибка при плавном завершении сервера", util.Err(err))
		return
	}

	log.Info("Сервер успешно остановлен")
}
