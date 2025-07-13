package main

import (
    "log"
    "net/http"

    "segment-service/internal/config"
    "segment-service/internal/router"
	"segment-service/internal/db"

)

func main() {
    cfg := config.Load()

    db.InitDB() // подключаем БД

    r := router.NewRouter()
    log.Printf("Server starting on port %s...", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}

