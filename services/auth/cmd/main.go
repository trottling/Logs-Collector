package main

import (
	"context"
	"errors"
	"log"
	"logs-collector/services/auth/internal/api/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"logs-collector/services/auth/internal/config"
	"logs-collector/services/auth/internal/jwt"
	"logs-collector/services/auth/internal/store"
)

func main() {
	// Загрузка конфига
	cfg := config.Load()

	// GORM подключение
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Мигрирование и загрузка расширения pg_crypto
	st := store.New(db)
	if err := st.Migrate(context.Background()); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// Создание root юзера при отсутствии
	if err := st.SeedRoot(cfg.RootLogin, cfg.RootPassword); err != nil {
		log.Fatalf("seed root: %v", err)
	}

	jm := jwt.New(cfg.JWTSecret)

	// Роутер
	handler := api.NewRouter(st, jm)

	// Настройки сервера
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// ✨ Graceful shutdown ✨
	go func() {
		log.Printf("auth listening on %s", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %v", err)
		}
	}()

	// ✨ Wait for stop ✨
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown: %v", err)
	}
	log.Println("server stopped")
}
